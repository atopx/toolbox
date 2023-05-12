package create

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/public_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.Name == "" {
		return nil, ecode.ECode_LABEL_PARAMS_ERROR_NameRequired
	}
	_, code := public_client.OperateLabel(c.Context(), &public_service.OperateLabelParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_CREATE,
		Data:    &public_service.Label{Name: params.Name},
	})
	return Reply{}, code
}
