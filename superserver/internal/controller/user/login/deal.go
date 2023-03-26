package login

import (
	"superserver/common/system"
	"superserver/common/utils"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
	"time"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.Username == "" {
		return nil, ecode.ECode_USER_PARAMS_ERROR_UsernameRequired
	}
	if params.Password == "" {
		return nil, ecode.ECode_USER_PARAMS_ERROR_PasswordRequired
	}
	listUserReply, code := auth_client.ListUser(c.Context, &auth_service.ListUserParams{
		Header: system.NewServiceHeader(c.Header),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &auth_service.UserFilter{Usernames: []string{params.Username}},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	if len(listUserReply.Data) == 0 {
		return nil, ecode.ECode_USER_PARAMS_ERROR_UserNotFound
	}

	user := listUserReply.Data[0]
	if utils.Hash(params.Password) != user.Password {
		return nil, ecode.ECode_USER_PARAMS_ERROR_PasswordInvalid
	}

	listTokenReply, code := auth_client.ListAuthToken(c.Context, &auth_service.ListAuthTokenParams{
		Header: system.NewServiceHeader(c.Header),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &auth_service.AuthTokenFilter{UserIds: []int32{user.Id}},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	current := time.Now().Local()

	var token *auth_service.AuthToken

	switch len(listTokenReply.Data) {
	case 0:
		token = new(auth_service.AuthToken)
	default:
		token = listTokenReply.Data[0]
		// 已有token且未过期，直接返回给用户
		if token.ExpireTime < current.Unix() {
			return listTokenReply.Data[0], ecode.ECode_SUCCESS
		}
	}

	// 没有token或已过期， 需要重新签发
	expires := current.Add(24 * time.Hour)
	token.UserId = user.Id
	token.IssuedTime = current.Unix()
	token.ExpireTime = expires.Unix()
	token.AccessToken = utils.SignToken(current, expires, token.ExpireTime)
	// 使用AccessToken的过期时间加密
	token.RefreshToken = utils.SignToken(current, current.Add(7*24*time.Hour), token.ExpireTime)

	operateTokenReply, code := auth_client.OperateAuthToken(c.Context, &auth_service.OperateAuthTokenParams{
		Header:  system.NewServiceHeader(c.Header),
		Operate: common.Operation_OPERATION_UPSERT,
		Data:    token,
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	return operateTokenReply.Data, ecode.ECode_SUCCESS
}
