package middleware

import (
	"cloudos/common/interface/common_iface"
	"cloudos/common/interface/ecode_iface"
	"cloudos/common/logger"
	"cloudos/common/system"
	"cloudos/internal/model/access"
	"cloudos/internal/model/permission"
	"cloudos/internal/model/user"
	"cloudos/internal/model/user_token"
	"errors"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		chain := system.GetChainMessage(ctx)

		// 查询access是否是游客权限
		accessPo, err := access.NewDao(m.db).Get(ctx.Request.Method, ctx.Request.RequestURI)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			chain.Message = "404"
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
			return
		}
		if err != nil {
			logger.Error(ctx, "auth accessPo.Get failed", zap.Error(err))
			chain.Message = "[14101]系统错误, 请联系管理员"
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, chain)
			return
		}

		if accessPo.Status != common_iface.AccessStatus_ACCESS_ANONYMOUS {

			// 不是游客权限, 进行认证
			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				chain.Message = "账户未登录"
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
				return
			}

			// token解析
			tokenStr := token[strings.Index(token, " ")+1:]
			claims, err := user_token.NewDao(m.db).First(func(tx *gorm.DB) *gorm.DB { tx.Where("access_token=?", tokenStr); return tx })
			if err != nil {
				logger.Error(ctx, "auth system.UnSignClaims failed", zap.Error(err))
				chain.Message = ecode_iface.ECode_AUTH_INVALID.String()
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
				return
			}

			// 过期时间验证
			if claims.ExpireTime < time.Now().Unix() {
				chain.Message = ecode_iface.ECode_AUTH_EXPIRED.String()
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
				return
			}

			// 非系统管理员进行权限验证
			if claims.UserId != user.SystemUser.Id {
				if !permission.NewDao(m.db).Inspector(claims.UserId, accessPo.Id) {
					chain.Message = "无权限"
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
					return
				}
			}

			chain.UserId = claims.UserId
		}

		ctx.Next()
	}
}
