package create

import (
	"superserver/common/utils"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.Name == "" {
		return nil, ecode.ECode_USER_PARAMS_ERROR_NameRequired
	}
	if params.Username == "" {
		return nil, ecode.ECode_USER_PARAMS_ERROR_UsernameRequired
	}
	if params.Password == "" {
		return nil, ecode.ECode_USER_PARAMS_ERROR_PasswordRequired
	}

	listReply, code := auth_client.ListUser(c.Context(), &auth_service.ListUserParams{
		Header: c.NewServiceHeader(),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &auth_service.UserFilter{Usernames: []string{params.Username}},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	if len(listReply.Data) > 0 {
		return nil, ecode.ECode_USER_PARAMS_ERROR_UsernameExist
	}

	operateReply, code := auth_client.OperateUser(c.Context(), &auth_service.OperateUserParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_CREATE,
		Data: &auth_service.User{
			Username: params.Username,
			Password: utils.Hash(params.Password),
			Status:   auth_service.UserStatus_USER_NORMAL,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	return operateReply.Data, ecode.ECode_SUCCESS
}
