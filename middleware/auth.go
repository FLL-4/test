package middleware

import (
	"net/http"
	"strings"

	"github.com/fangyanlin/gin-gorm-app/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件（示例，需要配合JWT使用）
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		
		if authHeader == "" {
			utils.UnauthorizedResponse(c, "Authorization header is required")
			c.Abort()
			return
		}
		
		// 验证token格式 "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.UnauthorizedResponse(c, "Invalid authorization header format")
			c.Abort()
			return
		}
		
		token := parts[1]
		
		// TODO: 在这里添加JWT验证逻辑
		// 示例：验证token是否有效
		if token == "" {
			utils.UnauthorizedResponse(c, "Invalid token")
			c.Abort()
			return
		}
		
		// 将用户信息存储到上下文中
		// c.Set("user_id", userID)
		
		c.Next()
	}
}

// AdminMiddleware 管理员权限中间件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文获取用户信息
		// userID, exists := c.Get("user_id")
		// if !exists {
		// 	utils.UnauthorizedResponse(c, "Unauthorized")
		// 	c.Abort()
		// 	return
		// }
		
		// TODO: 检查用户是否为管理员
		// isAdmin := checkIsAdmin(userID)
		// if !isAdmin {
		// 	utils.ForbiddenResponse(c, "Admin access required")
		// 	c.Abort()
		// 	return
		// }
		
		c.Next()
	}
}

// RateLimitMiddleware 限流中间件（简单示例）
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 实现限流逻辑
		// 可以使用 redis 或内存存储来实现
		c.Next()
	}
}
