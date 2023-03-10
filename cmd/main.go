package main

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/XSAM/otelsql"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/sqlx"
	"github.com/tsingsun/woocoo/web"
	webhander "github.com/tsingsun/woocoo/web/handler"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/graph"
	"github.com/woocoos/adminx/security"
	"github.com/woocoos/adminx/service/resource"
	"github.com/woocoos/adminx/service/snowflake"
	"github.com/woocoos/adminx/service/ucenter"
	"go.opentelemetry.io/contrib/propagators/b3"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/adminx/ent/runtime"
)

var (
	portalClient  *ent.Client
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
		web.RegisterMiddleware(authz.New()),
	)
	webSrv.Router().NoRoute()
	client, err := open(conf.Global(), "store.portal")
	if err != nil {
		log.Fatal(err)
	}
	// ??????????????????gqlserver????????????,??????????????????????????????
	gqlSrv, err := gql.RegisterSchema(webSrv,
		graph.NewSchema(
			graph.RegisterEntClient(client),
			graph.RegistryService(&ucenter.Service{Client: client}, &resource.Service{Client: client}),
		))
	if err != nil {
		log.Fatal(err)
	}
	// gqlserver??????????????????
	for _, srv := range gqlSrv {
		if entCacheLevel == 1 {
			// ????????????http request?????????
			useContextCache(srv)
		}
		// mutation??????
		srv.Use(entgql.Transactioner{TxOpener: client})
		// ????????????
		srv.AddTransport(&transport.Websocket{})
	}
	// ????????????,??????????????????????????????????????????.
	if jwt, ok := webSrv.HandlerManager().Get("jwt"); ok {
		if h, ok := jwt.(*webhander.JWTMiddleware); ok {
			h.Config.ErrorHandler = func(c *gin.Context, err error) error {
				if c.IsWebsocket() && cnf.Development {
					return nil
				}
				return err
			}
		}
	}

	return webSrv
}

// ??????:
//  1. otelsql.AllowRoot()????????????,????????????????????????(warmup,????????????,????????????????????????)
//  2. ??????????????????????????????,????????????context.Background()??????
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
	// ??????????????????????????????,????????????entcache.ContextLevel()??????
	switch entCacheLevel {
	case 1:
		// ??????Context??????,????????????????????????ttl
		return entcache.NewDriver(drvori, entcache.ContextLevel())
	case 2:
		// ??????db??????
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
	err := security.SetAuthorization(cnf.Sub("authz"), drv)
	if err != nil {
		log.Fatal(err)
	}
}

func useContextCache(server *handler.Server) {
	server.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		if op := graphql.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}
		return next(ctx)
	})
}
