package xredis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
	"time"
)

func NewClient(key string) (client *redis.Client, err error) {
	if !viper.IsSet(key) {
		err = errors.New(fmt.Sprintf("redis config nil: %s", key))
		return
	}
	client = redis.NewClient(
		&redis.Options{
			Addr:         viper.GetString(key + ".addr"),
			Password:     viper.GetString(key + ".password"),
			DB:           viper.GetInt(key + ".db"),
			DialTimeout:  viper.GetDuration(key+".dialTimeout") * time.Millisecond,
			ReadTimeout:  viper.GetDuration(key+".readTimeout") * time.Millisecond,
			WriteTimeout: viper.GetDuration(key+".writeTimeout") * time.Millisecond,
			MaxRetries:   viper.GetInt(key + ".maxRetries"),
			PoolSize:     viper.GetInt(key + ".poolSize"),
			MinIdleConns: viper.GetInt(key + ".minIdleConns"),
		})
	err = client.Ping().Err()
	if err != nil {
		log.Println("redis ping error",err)
	}
	return
}