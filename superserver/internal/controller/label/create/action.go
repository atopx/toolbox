package create

import (
	"fmt"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/public_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	fmt.Printf("%+v\n", c.Params)
	params := c.Params.(*Params)
	if _, ok := common.Source_name[params.Source]; !ok {
		return nil, ecode.ECode_SOURCE_PARAMS_ERROR_SourceInvalid
	}
	if params.Name == "" {
		return nil, ecode.ECode_LABEL_PARAMS_ERROR_NameRequired
	}
	_, code := public_client.OperateLabel(c.Context(), &public_service.OperateLabelParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_CREATE,
		Data: &public_service.Label{
			Source: params.Source,
			Name:   params.Name,
		},
	})
	return Reply{}, code
}
