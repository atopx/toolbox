package model

import (
	"superserver/common/utils"
	"superserver/domain/note_service"
)

type Note struct {
	Id         int32    `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	Labels     []string `json:"labels"`
	Topic      []string `json:"topics"`
	CreateTime string   `json:"create_time"`
	UpdateTime string   `json:"update_time"`
}

func NewNote(src *note_service.Note) Note {
	return Note{
		Id:         src.Id,
		Title:      src.Title,
		Content:    src.Content,
		Labels:     []string{},
		Topic:      []string{},
		CreateTime: utils.TimestampDefaultDecoder(src.CreateTime),
		UpdateTime: utils.TimestampDefaultDecoder(src.UpdateTime),
	}
}

func NewNoteList(list []*note_service.Note) []Note {
	result := make([]Note, 0, len(list))
	for _, src := range list {
		src.Content = "" // 列表不需要返回内容，减少数据量
		result = append(result, NewNote(src))
	}
	return result
}
