package middleware

import (
	"cloudos/common/interface/ecode_iface"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"cloudos/common/logger"
	"cloudos/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoverMiddleware 崩溃恢复中间件
func (m *Middleware) RecoverMiddleware() gin.HandlerFunc {
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
				stack := zap.Stack("stack")
				logger.Error(ctx, "recovery from panic",
					zap.Any("error", err),
					zap.ByteString("body", body),
					stack,
				)
				log.Println(stack.String)
				chain := system.GetChainMessage(ctx)
				chain.WriteAbnormal(ecode_iface.ECode_SYSTEM_ERROR, http.StatusText(http.StatusInternalServerError))
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, chain)
			}
		}()
		ctx.Next()
	}
}
