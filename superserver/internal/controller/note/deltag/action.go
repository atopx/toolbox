package deltag

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
	var labelType note_service.LabelType
	if params.IsTopic {
		labelType = note_service.LabelType_TOPIC
	}

	listNoteLabelReply, code := note_client.ListNoteLabel(c.Context(), &note_service.ListNoteLabelParams{
		Header: c.NewServiceHeader(),
		Pager:  &common.Pager{Disabled: true},
		Filter: &note_service.NoteLabelFilter{
			NoteIds:    []int32{params.NoteId},
			LabelIds:   []int32{params.LabelId},
			LabelTypes: []note_service.LabelType{labelType},
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	if len(listNoteLabelReply.Data) == 0 {
		return nil, ecode.ECode_BAD_REQUEST
	}

	_, code = note_client.OperateNoteLabel(c.Context(), &note_service.OperateNoteLabelParams{
		Header:  c.NewServiceHeader(),
		Operate: common.Operation_OPERATION_DELETE,
		Data:    &note_service.NoteLabel{Id: listNoteLabelReply.Data[0].Id},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	return nil, ecode.ECode_SUCCESS
}
