package delete

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/public_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	_, code := public_client.OperateFolder(c.Context(), &public_service.OperateFolderParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_DELETE,
		Data:    &public_service.Folder{Id: params.Id},
	})

	return Reply{}, code
}
