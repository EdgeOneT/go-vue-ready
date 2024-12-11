package services

import (
	"github.com/mojocn/base64Captcha"
	"go-vue-ready/day7-bse64captcha/database"
	"golang.org/x/net/context"
	"time"
)

type CaptchaService struct{}

// store保存生成的CAPTCHA数据
var store = base64Captcha.DefaultMemStore

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{}
}

// GenerateCaptcha 生成验证码
func (svc *CaptchaService) GenerateCaptcha() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取验证码 id, b64s 图片, 和正确答案
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return "", "", err
	}
	// 可选：将验证码答案存储到 Redis，缓存时间 5 分钟
	// 缓存 id 和答案到 Redis
	database.Rdb.Set(context.Background(), id, answer, time.Minute*5)

	return id, b64s, nil
}

// VerifyCaptcha 校验验证码
func (svc *CaptchaService) VerifyCaptcha(id, value string) bool {
	cached, err := database.Rdb.Get(context.Background(), id).Result()
	if err != nil {
		return false
	}
	return cached == value
}
