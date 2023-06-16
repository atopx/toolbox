package middleware

import (
	"cloudos/common/pb"
	"cloudos/common/system"
	"cloudos/internal/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		resp := system.GetResponse(ctx)

		defer func() {
			switch resp.Code {
			case pb.ECode_SUCCESS:
				ctx.Next()
			default:
				ctx.AbortWithStatusJSON(http.StatusOK, resp)
			}
		}()

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			resp.Code = pb.ECode_AuthTokenInvalid
			return
		}

		// token解析
		tokenStr := token[strings.Index(token, " ")+1:]

		// token不存在
		authToken := new(model.AuthTokenDao).First("access_token = ?", tokenStr)
		if authToken == nil {
			resp.Code = pb.ECode_NotFoundToken
			return
		}

		// token过期
		if authToken.ExpireTime < time.Now().Local().Unix() {
			resp.Code = pb.ECode_AuthTokenExpired
		}
	}
}
