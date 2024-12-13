package routes

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day7-bse64captcha/controllers"
)

func RegisterCaptchaRouter(r *gin.Engine, captchaController *controllers.CaptchaController) {
	r.GET("/captcha", captchaController.GenerateCaptcha)
	r.POST("/captcha/verify", captchaController.VerifyCaptcha)
}
