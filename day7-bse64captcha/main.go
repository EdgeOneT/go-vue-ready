package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-vue-ready/day7-bse64captcha/config"
	"go-vue-ready/day7-bse64captcha/controllers"
	"go-vue-ready/day7-bse64captcha/database"
	_ "go-vue-ready/day7-bse64captcha/docs"
	"go-vue-ready/day7-bse64captcha/repositories"
	"go-vue-ready/day7-bse64captcha/routes"
	"go-vue-ready/day7-bse64captcha/services"
)

func main() {
	// 1. 加载配置文件（数据库、端口等）
	config.LoadConfig("config.yaml")
	// 2. 初始化数据库连接
	database.InitDB()

	// 3. 实例化依赖
	//userRepo := repositories.NewUserRepository(database.DB)      // 创建用户数据存储实例
	//userService := services.NewUserService(userRepo)             // 创建用户服务实例
	//userController := controllers.NewUserController(userService) // 创建用户控制器实例

	// 4. 初始化 Gin 路由
	r := gin.Default()

	// 5. 注册用户路由
	routes.RegisterUserRoutes(r, controllers.NewUserController(services.NewUserService(repositories.NewUserRepository(database.DB))))
	routes.RegisterRoutes(r, controllers.NewRoleController(services.NewRoleService(repositories.NewRoleRepository(database.DB))))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 6. 启动服务，监听配置文件中指定的端口
	port := config.Cfg.Server.Port
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
