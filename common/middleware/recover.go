package middleware

import "github.com/gin-gonic/gin"

// RecoverMiddleware 崩溃恢复中间件 TODO 待实现
func RecoverMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
