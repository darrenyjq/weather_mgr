package internal

import (
	"base/pkg/xzap"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

var (
	CurrentEnv	=  os.Getenv("CUSTOM_RUNTIME_ENV")
	LocationEnv	=  os.Getenv("LOCATION")
)

func init()  {
	if CurrentEnv == "" || LocationEnv == "" {
		panic("环境变量为空，停止运行")
		return
	}
	log.Println("step 0: init start env:", CurrentEnv, LocationEnv)

	// init config 部署上线时，要制定下生产配置路径
	cfgFile := fmt.Sprintf("/go/src/withdraw/configs/%s/%s/conf.toml",strings.ToLower(CurrentEnv), strings.ToLower(LocationEnv))
	viper.SetConfigFile(cfgFile)
	err := viper.ReadInConfig()
	if err!=nil {
		panic(err)
	}
	log.Println("step 1: load config success", viper.ConfigFileUsed())

	//init zlog
	err = xzap.InitZLog(viper.GetStringSlice("log.outputPaths"), zapcore.Level(viper.GetInt("log.level")))
	if err!=nil {
		panic(err)
	}
}