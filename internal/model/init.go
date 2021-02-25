package model

import (
	"context"
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"weather_mgr/internal"
	"weather_mgr/pkg/xredis"
	"weather_mgr/pkg/xzap"
)

var (
	redisClient *redis.Client
)
var err error

func init() {
	// 初始化 redis
	redisAddr := internal.GetApolloCli().GetStringValue("redis.addr", "application", "")
	redisPWD := internal.GetApolloCli().GetStringValue("redis.password", "application", "")
	redisDB := internal.GetApolloCli().GetIntValue("redis.db", "application", 0)
	if redisClient, err = xredis.NewClient(redisAddr, redisPWD, redisDB); err != nil {
		xzap.Debug(err.Error())
		panic(err)
	}
}

func GetTimeNow(ctx context.Context) time.Time {
	millisecVal := ctx.Value("X-TS-Millisec")
	if millisecVal != nil {
		millisec, _ := millisecVal.(string)
		userTime, err := MsToTime(millisec)
		if err == nil {
			return userTime
		}
	}
	return time.Now()
}
func MsToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	tm := time.Unix(0, msInt*int64(time.Millisecond))
	return tm, nil
}
