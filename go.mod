module github.com/woocoos/adminx

go 1.18

require (
	ariga.io/entcache v0.1.1-0.20220825100256-1b8bbeb21c75
	entgo.io/contrib v0.3.5
	entgo.io/ent v0.11.8
	github.com/99designs/gqlgen v0.17.24
	github.com/XSAM/otelsql v0.18.0
	github.com/bwmarrin/snowflake v0.3.1-0.20221123153919-bc74ab286f15
	github.com/go-sql-driver/mysql v1.7.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/stretchr/testify v1.8.1
	github.com/tsingsun/woocoo v0.2.2-0.20230219115027-3872cb5faff9
	github.com/tsingsun/woocoo/contrib/telemetry v0.0.0-20230202042805-9108faa63c9d
	github.com/vektah/gqlparser/v2 v2.5.1
	github.com/vmihailenco/msgpack/v5 v5.3.5
	github.com/woocoos/casbin-ent-adapter v0.0.0-20230208061537-0374390b7bca
	go.opentelemetry.io/contrib/propagators/b3 v1.13.0
	go.opentelemetry.io/otel v1.12.0
	go.uber.org/zap v1.24.0
	golang.org/x/sync v0.1.0
)

require (
	ariga.io/atlas v0.9.1-0.20230119145809-92243f7c55cb // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/casbin/casbin/v2 v2.61.1 // indirect
	github.com/casbin/redis-watcher/v2 v2.3.0 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.8.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.13 // indirect
	github.com/knadh/koanf v1.4.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.38.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.12.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.35.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.12.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.12.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v0.35.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.12.0 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
	go.opentelemetry.io/otel/sdk v1.12.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.12.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20230131230820-1c016267d619 // indirect
	google.golang.org/grpc v1.52.3 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
