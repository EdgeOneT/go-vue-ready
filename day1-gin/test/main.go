package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		cur := time.Now()
		c.Header("X-Server-Time", cur.Format("2006-01-01 15:02:01"))
	}
}

func main() {
	r := gin.Default()
	r.Use(TimeMiddleware())
	reqJson := gin.H{}
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.POST("/echo", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&reqJson); err != nil {
			c.JSON(http.StatusBadRequest, "invalid JSON")
		}
		c.JSON(200, reqJson)
	})
	r.Run(":8080")
}
