package xconfig

import (
	"context"
	"encoding/json"
	"weather_mgr/pkg/xzap"

	"go.uber.org/zap"
)

func GetCoinAwardLimitConfig(ctx context.Context, appName string) (config map[string]int64, err error) {
	config = map[string]int64{}
	jsonByte, err := cfg.getConfig(ctx, "coin_award_limit", 1, appName, false)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonByte, &config)
	if err != nil {
		xzap.ErrorContext(ctx, "coin_award_limit", zap.Error(err))
	}
	return
}
