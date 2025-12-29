package middleware
package middleware

import (
	"net/http"

	"github.com/fangyanlin/gin-gorm-app/utils"
	"github.com/gin-gonic/gin"
)

// Recovery 错误恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误日志
				// log.Printf("Panic recovered: %v", err)
				
				// 返回500错误
				utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
				c.Abort()
			}
		}()
		
		c.Next()
	}
}
