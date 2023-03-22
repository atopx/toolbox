package role

import (
	"superserver/common/interface/user_iface"
	"superserver/internal/model"
)

type Role struct {
	model.Base

	Name    string
	Nature  user_iface.RoleNature
	Creator int // 创建人
	Updater int // 更新人
}

func New(id int) *Role {
	return &Role{Base: model.Base{Id: id}}
}

var (
	SystemRole  *Role
	DefaultRole *Role
)
