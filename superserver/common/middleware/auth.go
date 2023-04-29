package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"superserver/common/system"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
	"time"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		header := system.GetRequestHeader(ctx)
		responseHeader := system.GetResponseHeader(ctx)
		serviceHeader := system.NewServiceHeader(header)
		servicePager := &common.Pager{Index: 1, Size: 1}

		defer func() {
			if responseHeader.Code != ecode.ECode_SUCCESS {
				ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
			} else {
				ctx.Next()
			}
		}()

		// 查询access是否是游客权限
		listAccessReply, code := auth_client.ListAccess(ctx, &auth_service.ListAccessParams{
			Header: system.NewServiceHeader(header),
			Pager:  servicePager,
			Filter: &auth_service.AccessFilter{
				Methods: []string{ctx.Request.Method},
				Paths:   []string{ctx.Request.RequestURI},
			},
		})
		if code != ecode.ECode_SUCCESS {
			responseHeader.Code = code
			return
		}

		if len(listAccessReply.Data) == 0 {
			responseHeader.Code = ecode.ECode_ACCESS_NotFound
			return
		}

		accessPo := listAccessReply.Data[0]

		switch accessPo.Status {

		case auth_service.AccessStatus_ACCESS_ANONYMOUS: // 不处理, 直接放行

		case auth_service.AccessStatus_ACCESS_DISABLED:
			// 禁止访问
			responseHeader.Code = ecode.ECode_ACCESS_Disabled
			return

		default: // 默认权限控制
			// 获取token
			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				responseHeader.Code = ecode.ECode_AUTH_TOKEN_Required
				return
			}

			// token解析
			tokenStr := token[strings.Index(token, " ")+1:]
			listAuthTokenReply, code := auth_client.ListAuthToken(ctx, &auth_service.ListAuthTokenParams{
				Header: serviceHeader,
				Pager:  &common.Pager{Index: 1, Size: 1},
				Filter: &auth_service.AuthTokenFilter{AccessTokens: []string{tokenStr}},
			})
			if code != ecode.ECode_SUCCESS {
				responseHeader.Code = code
				return
			}

			if len(listAuthTokenReply.Data) == 0 {
				responseHeader.Code = ecode.ECode_AUTH_TOKEN_NotFound
				return
			}

			authToken := listAuthTokenReply.Data[0]

			if authToken.ExpireTime < time.Now().Local().Unix() {
				responseHeader.Code = ecode.ECode_AUTH_TOKEN_Expired
				return
			}

			// set operator
			header.Operator = authToken.UserId

			// is admin user?
			if authToken.UserId != system.User.Id {
				listUserRoleRefReply, code := auth_client.ListUserRoleRef(ctx, &auth_service.ListUserRoleRefParams{
					Header: serviceHeader,
					Pager:  &common.Pager{Disabled: true},
					Filter: &auth_service.UserRoleRefFilter{
						UserIds: []int32{authToken.UserId},
					},
				})
				if code != ecode.ECode_SUCCESS {
					responseHeader.Code = code
					return
				}
				role := listUserRoleRefReply.Data[0]

				listPermissionReply, code := auth_client.ListPermission(ctx, &auth_service.ListPermissionParams{
					Header: serviceHeader,
					Pager:  servicePager,
					Filter: &auth_service.PermissionFilter{
						AccessIds: []int32{accessPo.Id},
						RoleIds:   []int32{role.Id},
					},
				})
				if code != ecode.ECode_SUCCESS {
					responseHeader.Code = code
					return
				}
				if len(listPermissionReply.Data) == 0 {
					responseHeader.Code = ecode.ECode_ACCESS_Forbidden
					return
				}
			}
		}
	}
}
