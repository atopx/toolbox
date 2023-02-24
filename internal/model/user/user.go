package user

import (
	"crypto/md5"
	"encoding/hex"
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

func (m *User) GetId() any {
	return m.Id
}

func (u *User) SetPassword(password string) {
	u.Password = u.Crypto(password)
}

func (u *User) ComparePassword(password string) bool {
	return u.Password == u.Crypto(password)
}

func (*User) Crypto(password string) string {
	if password == "" {
		return password
	}
	buf := []byte(password)
	m := md5.New()
	m.Write(buf)
	return hex.EncodeToString(m.Sum(buf[0:1]))
}
