package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		fmt.Printf("日志中间件：[%s] %s %s %s", method, path, clientIP, duration)
	}
}
