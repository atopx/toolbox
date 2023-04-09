package list

import (
	"superserver/common/system"
	"superserver/domain/note_service"
	"superserver/domain/public/ecode"
	"superserver/service/note_client"
)

func (c *Controller) Deal() (any, ecode.ECode) {
	params := c.Params.(*Params)

	if len(params.Filter.LabelIds) > 0 {
		// TODO 查笔记-标签关系 note_label
	}

	listNoteReply, code := note_client.ListNote(c.Context, &note_service.ListNoteParams{
		Header: system.NewServiceHeader(c.Header),
		Pager:  params.Pager,
		Sorts:  params.Sorts,
		Filter: &note_service.NoteFilter{
			FolderIds:    params.Filter.FolderIds,
			TopicIds:     params.Filter.TopicIds,
			PublicSelect: params.Filter.PublicSelect,
			Keywords:     &note_service.NoteFilter_Keywords{Keyword: params.Filter.Keyword},
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	reply := Reply{
		Pager: listNoteReply.Pager,
		Notes: listNoteReply.Data,
	}
	return reply, ecode.ECode_SUCCESS
}
