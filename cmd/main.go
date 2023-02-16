package main

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/XSAM/otelsql"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/authz"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/sqlx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/gql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/graph"
	"github.com/woocoos/adminx/service/resource"
	"github.com/woocoos/adminx/service/snowflake"
	"github.com/woocoos/adminx/service/ucenter"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	entadapterent "github.com/woocoos/casbin-ent-adapter/ent"
	"go.opentelemetry.io/contrib/propagators/b3"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/adminx/ent/runtime"
)

var (
	portalClient  *ent.Client
	casbinClient  *entadapterent.Client
	entCacheLevel = 0
	entCacheTTL   = time.Minute
)

func main() {
	app := woocoo.New()
	if app.AppConfiguration().IsSet("otel") {
		otelCnf := app.AppConfiguration().Sub("otel")
		otelcfg := telemetry.NewConfig(otelCnf,
			telemetry.WithPropagator(b3.New()),
		)
		defer otelcfg.Shutdown()
	}
	snowflake.SetDefaultNode(app.AppConfiguration())

	buildCashbin(app.AppConfiguration())

	webSrv := newWebServer(app.AppConfiguration())
	app.RegisterServer(webSrv)

	defer func() {
		portalClient.Close()
		casbinClient.Close()
	}()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func newWebServer(cnf *conf.AppConfiguration) *web.Server {
	entCacheLevel = cnf.Int("entcache.level")
	if cnf.Duration("entcache.ttl") > 0 {
		entCacheTTL = cnf.Duration("entcache.ttl")
	}
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	client, err := open(conf.Global(), "store.portal")
	if err != nil {
		log.Fatal(err)
	}
	// 如果需要设置gqlserver的中间件,可以使用第一个返回值
	gqlSrv, err := gql.RegisterSchema(webSrv,
		graph.NewSchema(
			graph.RegisterEntClient(client),
			graph.RegistryService(&ucenter.Service{Client: client}, &resource.Service{Client: client}),
		))
	if err != nil {
		log.Fatal(err)
	}
	// gqlserver的中间件处理
	for _, srv := range gqlSrv {
		if entCacheLevel == 1 {
			// 启用针对http request的缓存
			useContextCache(srv)
		}
		// mutation事务
		srv.Use(entgql.Transactioner{TxOpener: client})
	}

	return webSrv
}

// 注意:
//  1. otelsql.AllowRoot()目前关闭,不需要记录的场景(warmup,如果开启,会导致大量的记录)
//  2. 在不需要记录的场景下,可以使用context.Background()代替
func buildEntDriver(cnf *conf.AppConfiguration, storekey string) dialect.Driver {
	var err error
	storeCfg := cnf.Sub(storekey)
	driverName := storeCfg.String("driverName")
	if cnf.IsSet("otel") {
		// Register the otelsql wrapper for the provided postgres driver.
		driverName, err = otelsql.Register("mysql",
			otelsql.WithAttributes(semconv.DBSystemMySQL),
			otelsql.WithAttributes(semconv.DBNameKey.String(storekey)),
			otelsql.WithSpanOptions(otelsql.SpanOptions{
				DisableErrSkip:  true,
				OmitRows:        true,
				OmitConnPrepare: true,
			}),
		)
		if err != nil {
			panic(err)
		}
		storeCfg.Parser().Set("driverName", driverName)
	}
	db := sqlx.NewSqlDB(storeCfg)
	drvori := entsql.OpenDB(driverName, db)
	// 如果需要设置缓存级别,可以使用entcache.ContextLevel()设置
	switch entCacheLevel {
	case 1:
		// 使用Context缓存,但是不使用缓存的ttl
		return entcache.NewDriver(drvori, entcache.ContextLevel())
	case 2:
		// 使用db缓存
		return entcache.NewDriver(drvori, entcache.TTL(entCacheTTL))
	}
	return drvori
}

func open(cnf *conf.AppConfiguration, storekey string) (*ent.Client, error) {
	drv := buildEntDriver(cnf, storekey)
	var opts = []ent.Option{ent.Driver(drv)}
	if cnf.Development {
		opts = append(opts, ent.Debug())
	}
	portalClient = ent.NewClient(opts...)
	return portalClient, nil
}

func buildCashbin(cnf *conf.AppConfiguration) {
	drv := buildEntDriver(cnf, "store.portal")
	casbinClient = entadapterent.NewClient(entadapterent.Driver(drv))
	adp, err := entadapter.NewAdapterWithClient(casbinClient)
	if err != nil {
		log.Fatal(err)
	}
	err = casbinClient.Schema.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	authz.SetAdapter(adp)
}

func useContextCache(server *handler.Server) {
	server.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		if op := graphql.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}
		return next(ctx)
	})
}
