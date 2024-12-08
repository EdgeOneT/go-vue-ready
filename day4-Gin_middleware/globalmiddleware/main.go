package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前的处理
		startTime := time.Now()

		// 继续请求处理
		c.Next()

		// 请求后的处理
		duration := time.Since(startTime)
		status := c.Writer.Status()
		fmt.Printf("收集成功：[%s] %s %d %s\n", c.Request.Method, c.Request.URL.Path, status, duration)
	}

}

func main() {
	r := gin.Default()
	r.Use(LoggerMiddleware())
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
