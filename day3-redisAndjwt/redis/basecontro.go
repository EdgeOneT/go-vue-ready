package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost: 6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(ctx).Result() //ping一下，看连接成功没。
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Connected to Redis successfully!")
}

// 增
func setData(key, value string) {
	err := rdb.Set(ctx, key, value, time.Minute*10).Err() //第四个参数是设置过期时间
	if err != nil {
		log.Fatal("Failed to set data:", err)
	}
	fmt.Printf("Set key: %s, value: %s\n", key, value)
}

// 查
func getData(key string) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatal("Failed to get data:", err)
	}
	fmt.Printf("Get key: %s, value: %s\n", key, val)
}

// 将数据库查询结果缓存到 Redis 中以提高查询效率
func cacheUserData(userID, userName string) {
	userKey := "user:" + userID
	err := rdb.Set(ctx, userKey, userName, time.Hour).Err()
	if err != nil {
		log.Fatal("Failed to cache user data:", err)
	}
	fmt.Printf("Cached user data: %s -> %s\n", userID, userName)
}

func main() {
	initRedis()
	setData("username", "小明")
	setData("username", "小红")
	setData("username", "小天")
	getData("username")
}
