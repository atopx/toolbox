package list

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/public_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	listFolderReply, code := public_client.ListFolder(c.Context(), &public_service.ListFolderParams{
		Header: c.NewServiceHeader(),
		Pager:  params.Pager,
		Filter: &public_service.FolderFilter{
			Ids:       params.Filter.Ids,
			ParentIds: params.Filter.ParentIds,
			Sources:   params.Filter.Sources,
			Keywords:  &public_service.FolderFilter_Keywords{Name: params.Filter.Keyword},
			Deleted:   common.BooleanScope_BOOL_FALSE,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	reply := Reply{
		Pager: listFolderReply.Pager,
		List:  []*public_service.Folder{},
	}
	if listFolderReply.Data != nil {
		reply.List = listFolderReply.Data
	}
	return reply, ecode.ECode_SUCCESS
}
