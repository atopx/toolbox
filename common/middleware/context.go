package middleware

import (
	"net/http"
	"time"

	"superserver/common/logger"
	"superserver/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// set context
		traceMessage := system.NewChainContext()
		ctx.Set(system.ChainContextKey, traceMessage)

		// request
		logger.Info(ctx, "request", zap.String("path", ctx.Request.URL.Path), zap.String("client", ctx.ClientIP()))

		// 执行接口逻辑
		beginTime := time.Now()
		ctx.Next()
		cost := zap.String("cost", time.Since(beginTime).String())
		switch traceMessage.GetLevel() {
		case system.CHAIN_BAD:
			logger.Warn(ctx, "response", cost, zap.String("warn", traceMessage.Message))
		case system.CHAIN_ERROR:
			logger.Error(ctx, "response", cost, zap.String("error", traceMessage.Message))
		default:
			logger.Info(ctx, "response", cost)
		}
		traceMessage.Message = http.StatusText(ctx.Writer.Status())
	}
}
