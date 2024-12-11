package database

import (
	"github.com/go-redis/redis/v8"
	"go-vue-ready/day7-bse64captcha/config"
	"golang.org/x/net/context"
	"log"
)

var Rdb *redis.Client

func initRedis() {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Host,
		Password: config.Cfg.Redis.Password,
		DB:       0,
	})
	if _, err := Rdb.Ping(ctx).Result(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
