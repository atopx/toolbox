package save

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/note_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	if params.Title == "" {
		return nil, ecode.ECode_NOTE_PARAMS_ERROR_TitleInvalid
	}

	note := &note_service.Note{
		Title:   params.Title,
		Content: params.Content,
	}

	var operate common.Operation
	switch params.Id {
	case 0:
		operate = common.Operation_OPERATION_CREATE
	default:
		note.Id = params.Id
		operate = common.Operation_OPERATION_UPDATE
	}
	_, code := note_client.OperateNote(c.Context(), &note_service.OperateNoteParams{
		Header:  c.NewServiceHeader(),
		Operate: operate,
		Data:    note,
	})

	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	if len(params.Labels) > 0 {
		// TODO create label

		// TODO create note-label
	}

	return Reply{}, code
}
