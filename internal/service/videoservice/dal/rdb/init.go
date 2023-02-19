package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func Init() {
	ctx := context.Background()
	RDB = redis.NewClient(&redis.Options{
		Password: "cuidongxu189",
		Addr:     "106.14.160.121:4251",
		DB:       1,
		PoolSize: 5,
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic("redis启动失败，请检查")
	}
}
