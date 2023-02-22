package user

import (
	"fmt"
	"superserver/internal/model"
	"superserver/proto/user_iface"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserDao struct {
	model.BaseDao
}

func (*UserDao) Tablename() string {
	return "su_user"
}

func NewDao(db *gorm.DB) *UserDao {
	dao := &UserDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.Tablename = dao.Tablename()
	return dao
}

func (dao *UserDao) FilterBy(key string, value any) error {
	return dao.Connection().Where(fmt.Sprintf("%s = ?", key), value).Error
}

func (dao *UserDao) InitSystemUser() error {
	SystemUser = &User{
		Name:     "系统管理员",
		Username: viper.GetString("server.admin.user"),
		Password: viper.GetString("server.admin.pass"),
		Role:     user_iface.UserRole_USER_ROLE_SYSTEM,
		Status:   user_iface.UserStatus_USER_STATUS_NORMAL,
	}
	return dao.Connection().Where(User{Username: SystemUser.Username, Role: SystemUser.Role}).FirstOrCreate(&SystemUser).Error
}
