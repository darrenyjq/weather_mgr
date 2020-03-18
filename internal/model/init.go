package model

import (
	"base/pkg/xmysql"
	"base/pkg/xredis"
	"database/sql"
	"github.com/go-redis/redis"
)

var (
	dbHandle	*sql.DB
	redisClient *redis.Client

)

func init()  {
	var err error

	//init handle
	if dbHandle,err = xmysql.NewMysql("db.local");err != nil {
		panic(err)
	}
	if redisClient,err = xredis.NewClient("redis.local");err != nil {
		panic(err)
	}
}
