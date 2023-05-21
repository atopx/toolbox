package model

import (
	"superserver/common/utils"
	"superserver/domain/note_service"
	"superserver/domain/public_service"
)

type Note struct {
	Id         int32                   `json:"id"`
	Title      string                  `json:"title"`
	Content    string                  `json:"content"`
	Labels     []*public_service.Label `json:"labels"`
	Topic      []*public_service.Label `json:"topics"`
	CreateTime string                  `json:"create_time"`
	UpdateTime string                  `json:"update_time"`
}

func NewNote(src *note_service.Note, labels, topics []*public_service.Label) Note {
	if labels == nil {
		labels = make([]*public_service.Label, 0)
	}
	if topics == nil {
		topics = make([]*public_service.Label, 0)
	}
	return Note{
		Id:         src.Id,
		Title:      src.Title,
		Content:    src.Content,
		Labels:     labels,
		Topic:      topics,
		CreateTime: utils.TimestampDefaultDecoder(src.CreateTime),
		UpdateTime: utils.TimestampDefaultDecoder(src.UpdateTime),
	}
}
