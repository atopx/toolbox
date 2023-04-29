package create

import (
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/auth_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.Name == "" {
		return nil, ecode.ECode_ROLE_PARAMS_ERROR_NameInvalid
	}

	reply, code := auth_client.ListRole(c.Context(), &auth_service.ListRoleParams{
		Header: c.NewServiceHeader(),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &auth_service.RoleFilter{
			Names:   []string{params.Name},
			Natures: []auth_service.RoleNature{auth_service.RoleNature_ROLE_CUSTOM},
		},
	})

	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	if reply.Pager.Count > 0 {
		return nil, ecode.ECode_ROLE_PARAMS_ERROR_NameExist
	}

	_, code = auth_client.OperateRole(c.Context(), &auth_service.OperateRoleParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_CREATE,
		Data: &auth_service.Role{
			Name:   params.Name,
			Nature: auth_service.RoleNature_ROLE_CUSTOM,
		},
	})
	return Reply{}, code
}
