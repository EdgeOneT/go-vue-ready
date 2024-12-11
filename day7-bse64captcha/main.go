package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-vue-ready/day4-Gin_middleware/test2/config"
	"go-vue-ready/day5-corefunc/controllers"
	"go-vue-ready/day5-corefunc/database"
	"go-vue-ready/day5-corefunc/repositories"
	"go-vue-ready/day5-corefunc/routes"
	"go-vue-ready/day5-corefunc/services"
)

func main() {
	// 1. 加载配置文件（数据库、端口等）
	config.LoadConfig("config.yaml")
	// 2. 初始化数据库连接
	database.InitDB()

	// 3. 实例化依赖
	userRepo := repositories.NewUserRepository(database.DB)      // 创建用户数据存储实例
	userService := services.NewUserService(userRepo)             // 创建用户服务实例
	userController := controllers.NewUserController(userService) // 创建用户控制器实例

	// 4. 初始化 Gin 路由
	r := gin.Default()

	// 5. 注册用户路由
	routes.RegisterUserRoutes(r, userController)

	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 6. 启动服务，监听配置文件中指定的端口
	port := config.Cfg.Server.Port
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}