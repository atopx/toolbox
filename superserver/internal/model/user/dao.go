package user

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"superserver/internal/model"
	"superserver/proto/user_iface"
)

type UserDao struct {
	model.BaseDao
}

func (*UserDao) TableName() string {
	return "su_user"
}

func NewDao(db *gorm.DB) *UserDao {
	dao := &UserDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *UserDao) GetUserMapByIds(ids []int) (map[int]*User, error) {
	result := make(map[int]*User)
	if len(ids) == 0 {
		return result, nil
	}
	var users []*User
	err := dao.Connection().Where("id in ?", ids).Find(&users).Error
	if err != nil {
		return result, err
	}
	for _, user := range users {
		result[user.Id] = user
	}
	return result, nil
}

func (dao *UserDao) GetUserByUsername(username string, excludeDeleted bool) (*User, error) {
	tx := dao.Connection().Where("username = ?", username)
	if excludeDeleted {
		tx.Scopes(dao.NotDeleted)
	}
	var user User
	err := tx.First(&user).Error
	return &user, err
}

func (dao *UserDao) LoadSystemUser(roleId int) error {
	SystemUser = new(User)
	err := dao.Connection().Where("role_id=?", roleId).First(SystemUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		SystemUser.Name = "系统管理员"
		SystemUser.RoleId = roleId
		SystemUser.Status = user_iface.UserStatus_USER_NORMAL
		SystemUser.Username = viper.GetString("admin.user")
		SystemUser.SetPassword(viper.GetString("admin.pass"))
		err = dao.Create(SystemUser)
	}
	return err
}
