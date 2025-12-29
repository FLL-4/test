package main
package main

import (
	"log"

	"github.com/fangyanlin/gin-gorm-app/config"
	"github.com/fangyanlin/gin-gorm-app/database"
	"github.com/fangyanlin/gin-gorm-app/middleware"
	"github.com/fangyanlin/gin-gorm-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	
	// 初始化数据库
	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()
	
	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)
	
	// 创建路由
	router := gin.New()
	
	// 使用中间件
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())
	
	// 设置路由
	routes.SetupRoutes(router, database.GetDB())
	
	// 启动服务器
	addr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
