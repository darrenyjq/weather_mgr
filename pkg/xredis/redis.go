package xredis

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func NewClient(addr string, password string, db int) (client *redis.Client, err error) {
	client = redis.NewClient(
		&redis.Options{
			Addr:         addr,
			Password:     password,
			DB:           db,
			DialTimeout:  time.Duration(300) * time.Millisecond,
			ReadTimeout:  time.Duration(500) * time.Millisecond,
			WriteTimeout: time.Duration(500) * time.Millisecond,
			MaxRetries:   64,
			PoolSize:     10,
			MinIdleConns: 5,
		})
	err = client.Ping().Err()
	if err != nil {
		log.Println("redis ping error", err)
	}
	return
}
