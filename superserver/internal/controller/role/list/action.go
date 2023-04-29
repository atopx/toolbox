package list

import (
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/internal/model"
	"superserver/service/auth_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	listRoleReply, code := auth_client.ListRole(c.Context(), &auth_service.ListRoleParams{
		Header: c.NewServiceHeader(),
		Pager:  params.Pager,
		Filter: &auth_service.RoleFilter{
			Keywords: &auth_service.RoleFilter_Keywords{Name: params.Filter.Keyword},
			Deleted:  common.BooleanScope_BOOL_FALSE,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	reply := Reply{
		Pager: listRoleReply.Pager,
		List:  []model.Role{},
	}
	if listRoleReply.Data != nil {
		reply.List = model.NewRoleList(listRoleReply.Data)
	}
	return reply, ecode.ECode_SUCCESS
}
