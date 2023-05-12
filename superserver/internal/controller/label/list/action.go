package list

import (
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/public_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	listLabelReply, code := public_client.ListLabel(c.Context(), &public_service.ListLabelParams{
		Header: c.NewServiceHeader(),
		Pager:  params.Pager,
		Filter: &public_service.LabelFilter{
			Ids:      params.Filter.Ids,
			Keywords: &public_service.LabelFilter_Keywords{Keyword: params.Filter.Keyword},
			Deleted:  common.BooleanScope_BOOL_FALSE,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	reply := Reply{
		Pager: listLabelReply.Pager,
		List:  []*public_service.Label{},
	}
	if listLabelReply.Data != nil {
		reply.List = listLabelReply.Data
	}
	return reply, ecode.ECode_SUCCESS
}
