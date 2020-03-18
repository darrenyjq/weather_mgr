package xconfig

import (
	"context"
	"encoding/json"
)

func GetConfigWithdraw(ctx context.Context) (config WithdrawConfig,err error)  {
	config = WithdrawConfig{}
	jsonByte,err := cfg.getConfig(ctx, configMap[_CONFIG_KEY_WITHDRAW])
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonByte,&config)
	return
}