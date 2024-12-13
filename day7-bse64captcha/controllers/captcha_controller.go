package controllers

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day7-bse64captcha/services"
	"net/http"
)

type CaptchaController struct {
	Service *services.CaptchaService
}

func NewCaptchaController(service *services.CaptchaService) *CaptchaController {
	return &CaptchaController{service}
}

// 生成验证码
func (ctrl *CaptchaController) GenerateCaptcha(c *gin.Context) {
	id, b64s, err := ctrl.Service.GenerateCaptcha()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate captcha"})
		return
	}
	c.JSON(200, gin.H{"id": id, "captcha": b64s})
}

// 校验验证码
func (ctrl *CaptchaController) VerifyCaptcha(c *gin.Context) {
	var payload struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ctrl.Service.VerifyCaptcha(payload.ID, payload.Value) {
		c.JSON(http.StatusOK, gin.H{"message": "captcha verified"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid captcha"})
	}
}
