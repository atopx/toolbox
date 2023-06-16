package list

import (
	"cloudos/common/pb"
	"cloudos/internal/model"
	"time"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)

	dao := new(model.FolderDao)

	tx := dao.Db().Model(new(pb.Folder)).Scopes(dao.NotDeleted)
	tx.Where("parent_id = ?", params.ParentId)
	tx.Where("domain = ?", params.Domain)

	if params.Keyword != "" {
		tx.Where("name like ?", dao.Like(params.Keyword))
	}

	var folders []*pb.Folder
	if params.Pager != nil && !params.Pager.Disabled {
		tx.Count(&params.Pager.Count)
		folders = dao.Pager(tx, params.Pager)
	} else {
		tx.Find(&folders)
	}

	reply := Reply{
		Pager: params.Pager,
		List:  make([]Folder, 0, len(folders)),
	}

	for _, folder := range folders {
		reply.List = append(reply.List, Folder{
			Id:         folder.Id,
			Name:       folder.Name,
			CreateTime: time.Unix(folder.CreateTime, 0).Format(time.DateTime),
			UpdateTime: time.Unix(folder.UpdateTime, 0).Format(time.DateTime),
		})
	}

	return reply, pb.ECode_SUCCESS
}
