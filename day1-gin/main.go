package main

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day1-gin/controllers"
	"net/http"
)

func main() {
	r := gin.Default()

	// 定义简单路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello user!",
		})
	})
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"meaage": id,
		})
	})
	r.GET("/newuser/:id", controllers.GerUser)

	r.Run(": 8080")
}
