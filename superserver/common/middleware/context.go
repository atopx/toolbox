package middleware

import (
	"net/http"
	"superserver/common/interface/ecode_iface"
	"time"

	"superserver/common/logger"
	"superserver/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ContextMiddleware 日志中间件
func (m *Middleware) ContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// set context
		chain := system.NewChainMessage()
		chain.IntoContext(ctx)
		defer chain.Recycle()
		// request
		logger.Info(ctx, "request", zap.String("path", ctx.Request.URL.Path), zap.String("client", ctx.ClientIP()))

		// 执行接口逻辑
		beginTime := time.Now()
		ctx.Next()
		cost := zap.String("cost", time.Since(beginTime).String())

		if chain.Level < ecode_iface.ECode_BAD_REQUEST {
			logger.Info(ctx, "response", cost)
		} else if chain.Level < ecode_iface.ECode_SYSTEM_ERROR {
			logger.Warn(ctx, "response", cost, zap.String("warn", chain.Message))
		} else {
			logger.Error(ctx, "response", cost, zap.String("error", chain.Message))
		}
		chain.Message = http.StatusText(ctx.Writer.Status())
	}
}
