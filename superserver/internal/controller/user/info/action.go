package info

import (
	"superserver/common/system"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	if c.Header.Operator == 0 {
		return nil, ecode.ECode_USER_PARAMS_ERROR_UserNotFound
	}
	listUserReply, code := auth_client.ListUser(c.Context, &auth_service.ListUserParams{
		Header: system.NewServiceHeader(c.Header),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &auth_service.UserFilter{Ids: []int32{c.Header.Operator}},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	if len(listUserReply.Data) == 0 {
		return nil, ecode.ECode_USER_PARAMS_ERROR_UserNotFound
	}
	return listUserReply.Data[0], ecode.ECode_SUCCESS
}
