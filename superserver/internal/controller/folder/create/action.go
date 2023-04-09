package create

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/public_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if _, ok := common.Source_name[params.Source]; !ok {
		return nil, ecode.ECode_SOURCE_PARAMS_ERROR_SourceInvalid
	}
	if params.Name == "" {
		return nil, ecode.ECode_FOLDER_PARAMS_ERROR_NameRequired
	}
	_, code := public_client.OperateFolder(c.Context(), &public_service.OperateFolderParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_CREATE,
		Data: &public_service.Folder{
			ParentId: params.ParentId,
			Source:   params.Source,
			Name:     params.Name,
			Path:     params.Path,
		},
	})
	return Reply{}, code
}
