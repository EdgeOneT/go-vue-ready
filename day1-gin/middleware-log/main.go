package main

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day1-gin/controllers"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("[%s] %s %s %d %s",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
		)
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())
	r.GET("/logerhello/:id", controllers.GerUser)
	r.Run(":8080")
}
