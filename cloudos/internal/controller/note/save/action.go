package save

import (
	"cloudos/common/consts"
	"cloudos/common/pb"
	"cloudos/internal/model"
	"fmt"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)
	fmt.Println(params.HtmlContent)
	if params.Title == "" {
		return nil, pb.ECode_InvalidName
	}
	operator := c.UserId()
	dao := new(model.NoteDao)
	var note *pb.Note
	if params.Id > 0 {
		// update
		note = dao.First("id = ? and creator = ?", params.Id, operator)
		if note == nil {
			return nil, pb.ECode_RecordNotFound
		}
		note.Title = params.Title
		note.Topic = params.Topic
		note.Content = params.Content
		// clean old labels
		if err := dao.CleanLabels(note.Id); err != nil {
			return nil, pb.ECode_ServerInternalError
		}
	} else {
		// create
		note = &pb.Note{
			Title:   params.Title,
			Topic:   params.Topic,
			Content: params.Content,
			Creator: operator,
			Updater: operator,
		}
	}

	// FIXME: 需要事务, 回头优化
	if err := dao.Save(note); err != nil {
		return nil, pb.ECode_ServerInternalError
	}

	if err := dao.AddLabels(note.Id, params.Labels); err != nil {
		return nil, pb.ECode_ServerInternalError
	}

	return consts.EmptyObj, pb.ECode_SUCCESS
}
