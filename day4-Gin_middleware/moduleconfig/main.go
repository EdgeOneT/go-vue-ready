package main

import (
	"fmt"
	"go-vue-ready/day4-Gin_middleware/moduleconfig/config"
	"log"
)

func main() {
	// 加载配置
	config.LoadConfig("config.yaml")
	// 使用配置
	fmt.Printf("1:Server running on port: %d\n", config.Cfg.Server.Port)
	fmt.Printf("2:Database host: %s\n", config.Cfg.Database.Host)
	// 模拟使用 Redis 配置
	redisAddr := fmt.Sprintf("3:%s:%d", config.Cfg.Redis.Host, config.Cfg.Redis.Port)
	log.Printf("4:Connecting to Redis at %s\n", redisAddr)
}
