package delete

import (
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	_, code := auth_client.OperateRole(c.Context(), &auth_service.OperateRoleParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_DELETE,
		Data:    &auth_service.Role{Id: params.Id},
	})

	return Reply{}, code
}
