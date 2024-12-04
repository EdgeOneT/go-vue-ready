package controllers

import "github.com/gin-gonic/gin"

func GerUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"user_id": id,
		"name":    "test_name",
	})
}
