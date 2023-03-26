package middleware

import (
	"net/http"
	"strings"
	"superserver/common/system"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		header := system.GetRequestHeader(ctx)
		responseHeader := system.GetResponseHeader(ctx)
		serviceHeader := system.NewServiceHeader(header)
		servicePager := &common.Pager{Index: 1, Size: 1}

		// 查询access是否是游客权限
		listAccessReply, code := auth_client.ListAccess(ctx, &auth_service.ListAccessParams{
			Header: serviceHeader,
			Pager:  servicePager,
			Filter: &auth_service.AccessFilter{
				Methods: []string{ctx.Request.Method},
				Paths:   []string{ctx.Request.RequestURI},
			},
		})
		if code != ecode.ECode_SUCCESS {
			responseHeader.Code = code
			ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
			return
		}

		if len(listAccessReply.Data) == 0 {
			responseHeader.Code = ecode.ECode_ACCESS_NotFound
			ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
			return
		}

		accessPo := listAccessReply.Data[0]
		if accessPo.Status == auth_service.AccessStatus_ACCESS_DISABLED {
			responseHeader.Code = ecode.ECode_ACCESS_Disabled
			ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
			return
		}

		if accessPo.Status != auth_service.AccessStatus_ACCESS_ANONYMOUS {

			// 不是游客权限, 进行认证
			if token == "" || !strings.HasPrefix(token, "Bearer ") {
				responseHeader.Code = ecode.ECode_AUTH_TOKEN_Required
				ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
				return
			}

			// token解析
			tokenStr := token[strings.Index(token, " ")+1:]
			listAuthTokenReply, code := auth_client.ListAuthToken(ctx, &auth_service.ListAuthTokenParams{
				Header: serviceHeader,
				Pager:  servicePager,
				Filter: &auth_service.AuthTokenFilter{AccessTokens: []string{tokenStr}},
			})
			if code != ecode.ECode_SUCCESS {
				ctx.AbortWithStatusJSON(http.StatusOK, code)
				return
			}
			if len(listAuthTokenReply.Data) == 0 {
				responseHeader.Code = ecode.ECode_AUTH_TOKEN_NotFound
				ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
				return
			}

			authToken := listAuthTokenReply.Data[0]
			header.Operator = authToken.UserId

			if authToken.ExpireTime < time.Now().Local().Unix() {
				responseHeader.Code = ecode.ECode_AUTH_TOKEN_Expired
				ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
				return
			}

			if authToken.UserId != system.User.Id {
				listUserRoleRefReply, code := auth_client.ListUserRoleRef(ctx, &auth_service.ListUserRoleRefParams{
					Header: serviceHeader,
					Pager:  &common.Pager{Disabled: true},
					Filter: &auth_service.UserRoleRefFilter{
						UserIds: []int32{authToken.UserId},
					},
				})
				if code != ecode.ECode_SUCCESS {
					ctx.AbortWithStatusJSON(http.StatusOK, code)
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
					ctx.AbortWithStatusJSON(http.StatusOK, code)
					return
				}
				if len(listPermissionReply.Data) == 0 {
					responseHeader.Code = ecode.ECode_ACCESS_Forbidden
					ctx.AbortWithStatusJSON(http.StatusOK, system.NewErrorResponse(responseHeader))
					return
				}
			}
		}
		ctx.Next()
	}
}
