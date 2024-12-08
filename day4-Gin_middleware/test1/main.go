package main

import (
	"fmt"
	"go-vue-ready/day4-Gin_middleware/test1/config"
	"log"
)

func main() {
	config.LoadConfig("config.yaml")

	// 打印配置信息
	fmt.Printf("Server running on port: %d\n", config.Cfg.Server.Port)
	fmt.Printf("Database host: %s\n", config.Cfg.Database.Host)
	fmt.Printf("Redis host: %s:%d\n", config.Cfg.Redis.Host, config.Cfg.Redis.Port)

	log.Println("Application started...")
}
