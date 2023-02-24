package user

import (
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
	once.Do(func() { dao.init() })
	return dao
}

func (dao *UserDao) init() {
	SystemUser = &User{
		Name:     "系统管理员",
		Username: viper.GetString("admin.user"),
		Role:     user_iface.UserRole_USER_ROLE_SYSTEM,
		Status:   user_iface.UserStatus_USER_NORMAL,
	}
	SystemUser.SetPassword(viper.GetString("admin.pass"))
	err := dao.Connection().Where(User{Username: SystemUser.Username, Role: SystemUser.Role}).FirstOrCreate(&SystemUser).Error
	if err != nil {
		panic(err)
	}
}

func (dao *UserDao) Upsert(user *User) error {
	return dao.Connection().Where(User{Username: SystemUser.Username, Role: SystemUser.Role}).FirstOrCreate(&SystemUser).Error
}

func (dao *UserDao) GetUserMapByIds(ids []int) (map[int]*User, error) {
	result := make(map[int]*User)
	if len(ids) == 0 {
		return result, nil
	}
	var users []*User
	err := dao.Connection().Where("id in ?", ids).Find(users).Error
	if err != nil {
		return result, err
	}
	for _, user := range users {
		result[user.Id] = user
	}
	return result, nil
}
