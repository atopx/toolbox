package permission

import (
	"superserver/internal/model"
)

type Permission struct {
	model.Base
	AccessId int
	RoleId   int
	Creator  int64
	Updater  int64
}

func (m *Permission) GetId() any {
	return m.Id
}
