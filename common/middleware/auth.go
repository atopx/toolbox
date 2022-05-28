package middleware

import "github.com/gin-gonic/gin"

// AuthMiddleware 权限认证中间件 TODO 待实现
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
