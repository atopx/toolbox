package role

import (
	"superserver/internal/model"
	"superserver/proto/user_iface"
	"sync"
)

type Role struct {
	model.Base

	Name    string
	Nature  user_iface.RoleNature
	Creator int // 创建人
	Updater int // 更新人
}

func (m *Role) GetId() any {
	return m.Id
}

// RoleMap 角色缓存
var RoleMap sync.Map
var (
	SystemRole  *Role
	DefaultRole *Role
)
