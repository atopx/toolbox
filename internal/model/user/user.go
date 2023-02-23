package user

import (
	"superserver/internal/model"
	"superserver/proto/user_iface"
	"sync"
)

type User struct {
	model.Base
	Name     string                // 姓名
	Username string                // 用户名
	Password string                // 用户密码
	Role     user_iface.UserRole   // 用户角色
	Status   user_iface.UserStatus // 用户状态
}

var (
	once       sync.Once
	SystemUser *User
)
