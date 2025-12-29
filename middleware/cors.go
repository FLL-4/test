package middleware
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		
		if origin != "" {
			// 设置允许的来源
			c.Header("Access-Control-Allow-Origin", origin)
			// 设置允许的请求方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")




















































}	}		c.Next()				}			return			c.AbortWithStatus(204)		if method == "OPTIONS" {				}			c.Header("Access-Control-Allow-Credentials", "true")			c.Header("Access-Control-Allow-Headers", allowHeaders)			c.Header("Access-Control-Allow-Methods", allowMethods)						}				c.Header("Access-Control-Allow-Origin", origin)			if allowed {						}				}					break					allowed = true				if o == "*" || o == origin {			for _, o := range origins {			allowed := false			origins := strings.Split(allowOrigins, ",")			// 检查origin是否在允许列表中		if origin != "" {				origin := c.Request.Header.Get("Origin")		method := c.Request.Method	return func(c *gin.Context) {func CORSWithConfig(allowOrigins, allowMethods, allowHeaders string) gin.HandlerFunc {// CORSWithConfig 带配置的CORS中间件}	}		c.Next()				}			return			c.AbortWithStatus(204)		if method == "OPTIONS" {		// 放行所有OPTIONS方法				}			c.Header("Access-Control-Allow-Credentials", "true")			// 设置是否允许发送Cookie			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")			// 设置允许暴露的响应头			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")			// 设置允许的请求头