// 使用 Redis 存储已登录用户的 Token，模拟以下场景：
//  1. Token 的存储和验证。
//  2. Token 的过期时间控制。
//  3. 强制用户下线（删除 Token）。
package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

var secretKey = []byte("mysecretkey")

var ctx = context.Background()

var rdb *redis.Client

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// 生成 Token
func generateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // 1 小时过期
			Issuer:    "myApp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 验证 Token
func validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// Token存到Redis上
func storeTokenInRedis(username, token string) {
	cacheKey := "token:" + username
	err := rdb.Set(ctx, cacheKey, token, time.Hour).Err()
	if err != nil {
		panic(err)
	}
}

// 验证Token是不是存在Redis上
func validateTokenInRedis(username, token string) bool {
	cachekey := "token:" + username
	storedToken, err := rdb.Get(ctx, cachekey).Result()
	if err == redis.Nil {
		return false // Token 不存在
	} else if err != nil {
		panic(err) // Redis 错误
	}
	return storedToken == token
}

// 强制用户下线
func forceLogout(username string) {
	cacheKey := "token:" + username
	rdb.Del(ctx, cacheKey)
}

func main() {
	initRedis()
	r := gin.Default()

	// 登录接口
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username required"})
			return
		}

		token, err := generateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}

		// 存储到 Redis
		storeTokenInRedis(username, token)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// 验证接口
	r.GET("/validate", func(c *gin.Context) {
		username := c.Query("username")
		token := c.Query("token")

		if !validateTokenInRedis(username, token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "token is valid"})
	})

	// 强制下线接口
	r.POST("/logout", func(c *gin.Context) {
		username := c.PostForm("username")
		forceLogout(username)
		c.JSON(http.StatusOK, gin.H{"message": "user logged out"})
	})

	r.Run(":8080")
}
