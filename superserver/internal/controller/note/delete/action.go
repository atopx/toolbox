package delete

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/note_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.Id == nil {
		return nil, ecode.ECode_BAD_REQUEST
	}
	_, code := note_client.OperateNote(c.Context(), &note_service.OperateNoteParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_DELETE,
		Data:    &note_service.Note{Id: *params.Id},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	return nil, ecode.ECode_SUCCESS
}
