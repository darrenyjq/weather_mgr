package internal

import (
	"fmt"
	"weather_mgr/pkg/xzap"

	"gitlab.corp.cootek.com/cloud_infra/apollo-client-golang/apollo"
	"go.uber.org/zap/zapcore"
)

var apolloCli *apollo.ApolloClient

func GetApolloCli() *apollo.ApolloClient {
	if apolloCli == nil {
		apolloCli = apollo.NewCtkClient()
	}
	return apolloCli
}

func init() {
	//init zlog
	fmt.Println("start internal init")
	logPath := GetApolloCli().GetStringValue("log.path", "application", "/tmp/app.log")
	fmt.Println("logPath:", logPath)
	logLevel := GetApolloCli().GetIntValue("log.level", "application", -1)
	err := xzap.InitZLog([]string{logPath}, zapcore.Level(logLevel))
	if err != nil {
		panic(err)
	}
	xzap.Info("app init")
	fmt.Println("end internal init")
}
