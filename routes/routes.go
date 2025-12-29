package routes

import (
	"github.com/fangyanlin/gin-gorm-app/controller"
	"github.com/fangyanlin/gin-gorm-app/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 设置路由
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// 初始化控制器
	userController := controller.NewUserController(db)
	productController := controller.NewProductController(db)

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户路由
		users := v1.Group("/users")
		{
			users.POST("", userController.CreateUser)
			users.GET("", userController.GetUsers)
			users.GET("/search", userController.SearchUsers)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

		// 产品路由
		products := v1.Group("/products")
		{
			products.POST("", productController.CreateProduct)
			products.GET("", productController.GetProducts)
			products.GET("/search", productController.SearchProducts)
			products.GET("/category/:category", productController.GetProductsByCategory)
			products.GET("/:id", productController.GetProduct)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}
	}

	// 示例：使用认证中间件的路由组
	authenticated := v1.Group("/protected")
	authenticated.Use(middleware.AuthMiddleware())
	{
		authenticated.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "This is a protected route",
			})
		})
	}
}
