package model

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var (
	migRwDbHandle *sql.DB
	migRoDbHandle *sql.DB
	redisClient   *redis.Client
)

var (
	CoinModel *coinModel
)

func init() {
	// var err error
	// CoinModel = new(coinModel)
	// // 初始化 mysql
	// rwDBDsn := internal.GetApolloCli().GetStringValue("mig_rw_db.dsn", "application", "")
	// roDBDsn := internal.GetApolloCli().GetStringValue("mig_ro_db.dsn", "application", "")
	// //init handle
	// if migRwDbHandle, err = xmysql.NewMysql(rwDBDsn); err != nil {
	// 	panic(err)
	// }
	// if migRoDbHandle, err = xmysql.NewMysql(roDBDsn); err != nil {
	// 	panic(err)
	// }
	//
	// // 初始化 redis
	// redisAddr := internal.GetApolloCli().GetStringValue("redis.addr", "application", "")
	// redisPWD := internal.GetApolloCli().GetStringValue("redis.password", "application", "")
	// redisDB := internal.GetApolloCli().GetIntValue("redis.db", "application", 0)
	// if redisClient, err = xredis.NewClient(redisAddr, redisPWD, redisDB); err != nil {
	// 	log.Println("err:", err)
	// 	panic(err)
	// }
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
