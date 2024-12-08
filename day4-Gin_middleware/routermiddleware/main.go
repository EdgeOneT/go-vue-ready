package main

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "unauthorized"}) // 401 状态码：未授权（Unauthorized）
			c.Abort()                                   // 停止后续操作
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	authGroup := r.Group("/auth")
	authGroup.Use(AuthMiddleware()) // 为该路由组添加中间件
	{
		authGroup.GET("/secure", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized"})
		})
	}

	r.Run(":8080")
}
