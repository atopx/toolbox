package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"superserver/common/logger"
	"superserver/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoverMiddleware 崩溃恢复中间件
func RecoverMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				body, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logger.Error(ctx, "recovery from broken pipeline",
						zap.ByteString("body", body),
						zap.Any("error", err),
					)
					if err != nil {
						panic(err)
					}
					ctx.Abort()
					return
				}
				logger.Error(ctx, "recovery from panic",
					zap.Any("error", err),
					zap.ByteString("body", body),
					zap.Stack("stack"),
				)
				value, _ := ctx.Get(system.ChainContextKey)
				message := value.(*system.ChainContext)
				message.WriteAbnomal(system.CHAIN_ERROR, http.StatusText(http.StatusInternalServerError))
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, message)
			}
		}()
		ctx.Next()
	}
}
