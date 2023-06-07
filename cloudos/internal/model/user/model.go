package user

import (
	"cloudos/common/interface/user_iface"
	"cloudos/internal/model"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	model.Base
	Name     string                // 姓名
	Username string                // 用户名
	Password string                // 用户密码
	Status   user_iface.UserStatus // 用户状态
}

var SystemUser *User

func (m *User) SetPassword(password string) {
	m.Password = m.Crypto(password)
}

func (m *User) ComparePassword(password string) bool {
	return m.Password == m.Crypto(password)
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
