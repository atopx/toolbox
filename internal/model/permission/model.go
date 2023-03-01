package permission

import (
	"superserver/internal/model"
)

type Permission struct {
	model.Base
	AccessId int
	RoleId   int
	Creator  int
	Updater  int
}

func (m *Permission) GetId() any {
	return m.Id
}

func New(id int) *Permission {
	return &Permission{Base: model.Base{Id: id}}
}
