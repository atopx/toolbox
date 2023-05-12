package list

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/service/note_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	if len(params.Filter.LabelIds) > 0 {
		// TODO 查笔记-标签关系 note_label
	}

	listNoteReply, code := note_client.ListNote(c.Context(), &note_service.ListNoteParams{
		Header: c.NewServiceHeader(),
		Pager:  params.Pager,
		Sorts:  params.Sorts,
		Filter: &note_service.NoteFilter{
			Deleted:  common.BooleanScope_BOOL_FALSE,
			Keywords: &note_service.NoteFilter_Keywords{Keyword: params.Filter.Keyword},
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	reply := Reply{
		Pager: listNoteReply.Pager,
		List:  []*note_service.Note{},
	}
	if listNoteReply.Data != nil {
		reply.List = listNoteReply.Data
	}
	return reply, ecode.ECode_SUCCESS
}
