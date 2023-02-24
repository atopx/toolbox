package permission

import (
	"superserver/internal/model"
	"superserver/proto/user_iface"
)

type Permission struct {
	model.Base
	AccessId int
	RoleId   user_iface.UserRole
	Creator  int64
	Updater  int64
}

func (m *Permission) GetId() any {
	return m.Id
}
