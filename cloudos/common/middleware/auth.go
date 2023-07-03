package middleware

import (
	"cloudos/common/consts"
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

		resp := system.GetResponse(ctx)

		defer func() {
			switch resp.Code {
			case pb.ECode_SUCCESS:
				ctx.Next()
			default:
				resp.Message = resp.Code.String()
				ctx.AbortWithStatusJSON(http.StatusOK, resp)
			}
		}()

		role, ok := m.GetAccess(ctx)

		if !ok {
			resp.Code = pb.ECode_RecordNotFound
			return
		}

		if role == pb.UserRole_GUEST {
			return
		}

		token := ctx.Request.Header.Get("Authorization")

		if token == consts.EmptyStr || !strings.HasPrefix(token, "Bearer ") {
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
			return
		}

		user := new(model.UserDao).First("id = ?", authToken.UserId)

		if user == nil {
			resp.Code = pb.ECode_NotFoundUser
			return
		}

		if user.Status == pb.UserStatus_DISABLED.String() {
			resp.Code = pb.ECode_UserDisabled
			return
		}

		userRole := pb.UserRole(pb.UserRole_value[user.Role])
		if userRole < role {
			resp.Code = pb.ECode_Forbidden
			return
		}

		ctx.Set(consts.CK_UserId, user.Id)
		ctx.Set(consts.CK_UserRole, int(userRole))
	}
}
