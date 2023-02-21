package user

import (
	"crypto/md5"
	"encoding/hex"
	"superserver/common/utils"
	"superserver/internal/model"
)

type User struct {
	model.Base
	Username        string
	Password        string
	Role            string
	AvatarUrl       string
	Status          string
	TotalSizeLimit  int64
	SingleSizeLimit int64
	SizeUsed        int64
	LastLoginIp     int64
	LastLoginTime   int64
}

func (u *User) GetPk() any {
	return u.Id
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

func (u *User) SetLastLoginIp(ip string) {
	u.LastLoginIp = utils.IPEncode(ip)
}

func (u *User) GetLastLoginIp() string {
	return utils.IPDecode(u.LastLoginIp)
}
