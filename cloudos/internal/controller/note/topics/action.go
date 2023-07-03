package topics

import (
	"cloudos/common/pb"
	"cloudos/internal/model"
)

func (c *Controller) Deal() (any, pb.ECode) {
	dao := new(model.NoteDao)
	reply := new(Reply)
	dao.Db().Model(new(pb.Note)).
		Scopes(dao.NotDeleted).
		Where("creator = ?", c.UserId()).
		Group("topic").
		Select("topic").
		Scan(&reply)
	return reply, pb.ECode_SUCCESS
}
