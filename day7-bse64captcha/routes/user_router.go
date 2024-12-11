package routes

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day5-corefunc/controllers"
)

func RegisterUserRoutes(r *gin.Engine, userController *controllers.UserController) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/creat", userController.CreateUser)
		userGroup.GET("/get", userController.GetUser)
	}
}
