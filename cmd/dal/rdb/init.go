package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/okatu-loli/TikTokLite/config"
)

var RDB *redis.Client

func Init() {
	ctx := context.Background()
	RDB = redis.NewClient(&redis.Options{
		Password: config.RedisPasswordDEV,
		Addr:     config.RedisAddrDEV,
		DB:       1,
		PoolSize: 5,
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic("redis启动失败，请检查")
	}
}
