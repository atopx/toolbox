package update

import (
	"cloudos/common/consts"
	"cloudos/common/pb"
	"cloudos/internal/model"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)

	if params.Id <= 0 {
		return nil, pb.ECode_BadRequest
	}

	dao := new(model.FolderDao)
	folder := dao.First("id = ?", params.Id)
	if folder == nil {
		return nil, pb.ECode_NotFoundId
	}

	validObj := dao.First(
		"name=? and domain=? and parent_id=? and id!=?",
		params.Name, folder.Domain, params.ParentId, folder.Id,
	)

	if validObj != nil {
		return nil, pb.ECode_AlreadyExistName
	}

	folder.Name = params.Name
	folder.ParentId = params.ParentId

	if err := dao.Update(folder); err != nil {
		return nil, pb.ECode_ServerInternalError
	}

	return consts.EmptyObj, pb.ECode_SUCCESS
}
