package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{
				"user_id": id,
			})
		})
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "create user",
			})
		})
	}
	r.Run(":8080")
}
