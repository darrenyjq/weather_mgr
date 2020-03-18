package xconfig

import (
	"context"
	"fmt"
	"strings"
)

type configKey struct {
	key  string
	abFunc func(context.Context)(int64)
	objectStruct interface{}
}

func (this configKey) GetKey(ctx context.Context) (string) {
	return this.key
}

func (this configKey) GetGroup(ctx context.Context) (group int64){
	group = 1
	if this.abFunc != nil {
		group = this.abFunc(ctx)
	}
	return
}

func (this configKey) GetAppName(ctx context.Context) (AppName string) {
	AppName = "com.androidhealth.steps.money"
	val := ctx.Value("app_name")
	if val != nil {
		strVal,ok := val.(string)
		if ok && strVal != "" {
			AppName = strings.ToLower(strVal)
		}
	}
	return
}

func (this configKey) GetKeyGroupOSNameString(ctx context.Context) (string) {
	return fmt.Sprintf("%s-%d-%s", this.key, this.GetGroup(ctx), this.GetAppName(ctx))
}


func (this configKey) GetObjectStruct() interface{} {
	return this.objectStruct
}

//新增一个配置项
const (
	_CONFIG_KEY_WITHDRAW	=	"withdraw"

)
//新增配置项对应的内容
var configMap = map[string]configKey{
	_CONFIG_KEY_WITHDRAW	:	{_CONFIG_KEY_WITHDRAW,nil,WithdrawConfig{}},
}
