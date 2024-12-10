package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-vue-ready/day4-Gin_middleware/test2/middleware"

	"go-vue-ready/day4-Gin_middleware/test2/config"

	"net/http"
)

func main() {
	config.LoadConfig("config.yaml")
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.CORSMiddleware())
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username required"})
			return
		}
		token, err := middleware.GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token":    token,
			"username": username,
		})
	})
	// 受保护路由
	protected := r.Group("/secure")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/data", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Hello, %s!", username)})
	})
	port := config.Cfg.Server.Port
	r.Run(fmt.Sprintf(":%d", port))
}
