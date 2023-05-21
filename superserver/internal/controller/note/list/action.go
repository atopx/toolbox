package list

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/internal/model"
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
		Sorts:  []*common.Sort{{Field: "update_time", Direction: common.SortDirection_SORT_DESC}},
		Filter: &note_service.NoteFilter{
			Deleted:  common.BooleanScope_BOOL_FALSE,
			Keywords: &note_service.NoteFilter_Keywords{Keyword: params.Filter.Keyword},
		},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	var noteIds []int32
	for _, note := range listNoteReply.Data {
		noteIds = append(noteIds, note.Id)
	}

	noteLabelMap, code := note_client.GetNoteLabelMap(c, noteIds)
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	reply := Reply{
		Pager: listNoteReply.Pager,
		List:  make([]model.Note, 0, len(noteIds)),
	}

	for _, note := range listNoteReply.Data {
		labels := noteLabelMap[note.Id][note_service.LabelType_TAG]
		topics := noteLabelMap[note.Id][note_service.LabelType_TOPIC]
		reply.List = append(reply.List, model.NewNote(note, labels, topics))
	}

	return reply, ecode.ECode_SUCCESS
}
