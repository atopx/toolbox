package info

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/internal/model"
	"superserver/service/note_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)
	if params.Id == nil {
		return nil, ecode.ECode_BAD_REQUEST
	}
	listNoteReply, code := note_client.ListNote(c.Context(), &note_service.ListNoteParams{
		Header: c.NewServiceHeader(),
		Pager:  &common.Pager{Index: 1, Size: 1},
		Filter: &note_service.NoteFilter{
			Ids:     []int32{*params.Id},
			Deleted: common.BooleanScope_BOOL_FALSE,
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	if listNoteReply.Pager.Count != 1 {
		return nil, code
	}
	note := listNoteReply.Data[0]
	noteLabelMap, code := note_client.GetNoteLabelMap(c, []int32{note.Id})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	labels := noteLabelMap[note.Id][note_service.LabelType_TAG]
	topics := noteLabelMap[note.Id][note_service.LabelType_TOPIC]
	reply := Reply(model.NewNote(note, labels, topics))
	return reply, ecode.ECode_SUCCESS
}
