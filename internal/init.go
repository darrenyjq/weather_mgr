package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.corp.aaaa/cloud_infra/apollo-client-golang/apollo"
	"gitlab.corp.aaaa/cloud_infra/elete-go/pkg/elete/sdk"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"weather_mgr/pkg/xzap"
)

var apolloCli *apollo.ApolloClient

func GetApolloCli() *apollo.ApolloClient {
	if apolloCli == nil {
		apolloCli = apollo.NewCtkClient()
	}
	return apolloCli
}

func init() {
	// init zlog
	fmt.Println("start internal init")
	logPath := GetApolloCli().GetStringValue("log.path", "application", "/tmp/app.log")
	fmt.Println("logPath:", logPath)
	logLevel := GetApolloCli().GetIntValue("log.level", "application", -1)
	err := xzap.InitZLog([]string{"stderr", logPath}, zapcore.Level(logLevel))
	if err != nil {
		panic(err)
	}
	viper.Set("current_env", strings.ToLower(os.Getenv("CUSTOM_RUNTIME_ENV")))
	// 生产环境关闭 debug日志
	if viper.GetString("current_env") != "pro" {
		sdk.ChangeLogLevel(logLevel)
	}
	fmt.Println("end internal init")
}
