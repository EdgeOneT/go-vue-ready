package routes

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day7-bse64captcha/controllers"
)

func RegisterRoutes(r *gin.Engine, roleController *controllers.RoleController) {
	roleGroup := r.Group("/roles")
	{
		roleGroup.POST("", roleController.CreateRole)
		roleGroup.GET("", roleController.GetRoles)
		roleGroup.PUT("/:id", roleController.UpdateRole)
		roleGroup.DELETE("/id", roleController.DeleteRole)
	}
}
