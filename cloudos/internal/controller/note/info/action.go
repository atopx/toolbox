package info

import (
	"cloudos/common/pb"
	"cloudos/common/utils"
	"cloudos/internal/model"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)
	dao := new(model.NoteDao)
	note := dao.First("id=?", params.Id)
	if note == nil {
		return nil, pb.ECode_RecordNotFound
	}

	return Reply{
		Id:         note.Id,
		Title:      note.Title,
		Topic:      note.Topic,
		Content:    note.Content,
		Labels:     dao.Labels(note.Id),
		CreateTime: utils.Datetime(note.CreateTime),
		UpdateTime: utils.Datetime(note.UpdateTime),
	}, pb.ECode_SUCCESS
}
