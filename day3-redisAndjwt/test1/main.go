//将用户信息存储到 Redis 中，并模拟以下场景：
//	1. 首次从数据库获取用户信息并缓存到 Redis。
//	2. 缓存过期后重新从数据库获取数据。

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// 模拟数据库查询
func getUserFromDB(userID string) string {
	fmt.Printf("Fetching user %s from database...\n", userID)
	return fmt.Sprintf("User_%s_Data", userID) // 假设数据库返回的用户数据
}

// 获取用户信息
func getUserData(userID string) string {
	cacheKey := "user:" + userID
	data, err := rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		data = getUserFromDB(userID)
		rdb.Set(ctx, cacheKey, data, 10*time.Second)
	} else if err != nil {
		panic(err)
	}
	return data
}

func main() {
	initRedis()

	// 第一次获取数据
	userID := "123"
	fmt.Println("User Data:", getUserData(userID))

	// 模拟缓存失效
	fmt.Println("Waiting for cache to expire...")
	time.Sleep(12 * time.Second)

	// 再次获取数据，缓存失效后会重新加载
	fmt.Println("User Data:", getUserData(userID))
}
