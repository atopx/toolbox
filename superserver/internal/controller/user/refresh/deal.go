package refresh

import (
	"superserver/common/utils"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
	"time"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.RefreshToken == "" {
		return nil, ecode.ECode_AUTH_TOKEN_Required
	}
	listTokenReply, code := auth_client.ListAuthToken(c.Context(), &auth_service.ListAuthTokenParams{
		Header: c.NewServiceHeader(),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &auth_service.AuthTokenFilter{
			RefreshTokens: []string{params.RefreshToken},
			Deleted:       common.BooleanScope_BOOL_FALSE,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	if len(listTokenReply.Data) == 0 {
		return nil, ecode.ECode_AUTH_TOKEN_NotFound
	}
	token := listTokenReply.Data[0]

	current := time.Now().Local()
	// 如果已过期，重新签发
	expireDate, err := utils.UnSignTokenExpire(params.RefreshToken, token.ExpireTime)
	if err != nil {
		return nil, ecode.ECode_AUTH_TOKEN_Expired
	}
	if expireDate.Before(current) {
		expire := current.Add(24 * time.Hour)
		token.IssuedTime = current.Unix()
		token.ExpireTime = expire.Unix()
		// 使用AccessToken的过期时间加密
		token.AccessToken = utils.SignToken(current, expire, token.ExpireTime)
		token.RefreshToken = utils.SignToken(current, current.Add(7*24*time.Hour), token.ExpireTime)
		// update auth token
		if _, code = auth_client.OperateAuthToken(c.Context(), &auth_service.OperateAuthTokenParams{
			Header:  c.NewServiceHeader(),
			Operate: common.Operation_OPERATION_UPSERT,
			Data:    token,
		}); code != ecode.ECode_SUCCESS {
			return nil, code
		}
	}
	return token, ecode.ECode_SUCCESS
}
