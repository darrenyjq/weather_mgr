module base

go 1.13

require (
	cootek.com/elete/sdk v1.0.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/gavv/httpexpect/v2 v2.0.2
	github.com/gin-gonic/gin v1.4.0
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/onsi/ginkgo v1.10.2 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.1.0
	github.com/robfig/cron/v3 v3.0.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/uber-go/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	google.golang.org/grpc v1.21.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
)

replace (
	cootek.com/elete/sdk v1.0.0 => ../cootek.com/elete/sdk
	cootek.com/runtime v0.0.0 => ../cootek.com/runtime
)
