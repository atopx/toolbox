package create

import (
	"cloudos/common/consts"
	"cloudos/common/pb"
	"cloudos/internal/model"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)

	if params.Name == consts.EmptyStr {
		return nil, pb.ECode_InvalidName
	}

	if params.Domain != pb.DomainFolder_FILE.String() && params.Domain != pb.DomainFolder_NOTE.String() {
		return nil, pb.ECode_InvalidDomain
	}

	dao := new(model.FolderDao)

	folder := dao.First("parent_id=? and domain=?", params.ParentId, params.Domain)
	if folder != nil {
		return nil, pb.ECode_AlreadyExistName
	}

	folder = &pb.Folder{
		Name:     params.Name,
		ParentId: params.ParentId,
		Domain:   params.Domain,
	}

	if err := dao.Create(folder); err != nil {
		return nil, pb.ECode_ServerInternalError
	}

	return consts.EmptyObj, pb.ECode_SUCCESS
}
