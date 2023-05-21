package note_client

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service"
	"superserver/service/public_client"
)

func GetNoteLabelMap(bis service.Business, noteIds []int32) (map[int32]map[note_service.LabelType][]*public_service.Label, ecode.ECode) {
	result := make(map[int32]map[note_service.LabelType][]*public_service.Label)
	if len(noteIds) == 0 {
		return result, ecode.ECode_SUCCESS
	}
	// note-label ref
	listNoteLabelReply, code := ListNoteLabel(bis.Context(), &note_service.ListNoteLabelParams{
		Header: bis.NewServiceHeader(),
		Pager:  &common.Pager{Disabled: true},
		Filter: &note_service.NoteLabelFilter{NoteIds: noteIds},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}
	labelNum := len(listNoteLabelReply.Data)
	if labelNum == 0 {
		return result, ecode.ECode_SUCCESS
	}

	// list labels
	labelIds := make([]int32, 0, len(listNoteLabelReply.Data))
	noteLabelRef := make(map[int32]map[note_service.LabelType][]int32)
	for _, ref := range listNoteLabelReply.Data {
		labelIds = append(labelIds, ref.LabelId)
		if _, ok := noteLabelRef[ref.LabelId]; !ok {
			noteLabelRef[ref.LabelId] = make(map[note_service.LabelType][]int32)
		}
		noteLabelRef[ref.LabelId][ref.LabelType] = append(noteLabelRef[ref.LabelId][ref.LabelType], ref.NoteId)
	}
	listLabelReply, code := public_client.ListLabel(bis.Context(), &public_service.ListLabelParams{
		Header: bis.NewServiceHeader(),
		Pager:  &common.Pager{Disabled: true},
		Sorts:  []*common.Sort{{Field: "id", Direction: common.SortDirection_SORT_DESC}},
		Filter: &public_service.LabelFilter{Ids: labelIds},
	})
	if code != ecode.ECode_SUCCESS {
		return nil, code
	}

	// to map
	for _, label := range listLabelReply.Data {
		for t, refNoteIds := range noteLabelRef[label.Id] {
			for _, refNoteId := range refNoteIds {
				if _, ok := result[refNoteId]; !ok {
					result[refNoteId] = make(map[note_service.LabelType][]*public_service.Label)
				}
				result[refNoteId][t] = append(result[refNoteId][t], label)
			}
		}
	}
	return result, ecode.ECode_SUCCESS
}
