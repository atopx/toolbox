package addtag

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/note_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.NoteId == 0 || params.LabelId == 0 {
		return nil, ecode.ECode_BAD_REQUEST
	}
	// add note-label relation
	var labelType note_service.LabelType
	if params.IsTopic {
		labelType = note_service.LabelType_TOPIC
	}
	_, code := note_client.OperateNoteLabel(c.Context(), &note_service.OperateNoteLabelParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_CREATE,
		Data: &note_service.NoteLabel{
			NoteId:    params.NoteId,
			LabelId:   params.LabelId,
			LabelType: labelType,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	return nil, ecode.ECode_SUCCESS
}
