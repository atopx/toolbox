package user

import (
	"superserver/common/interface/user_iface"
	"superserver/internal/model"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_user"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *Dao) GetUserMapByIds(ids []int) (map[int]*User, error) {
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

func (dao *Dao) GetUserByUsername(username string, excludeDeleted bool) (*User, error) {
	tx := dao.Connection().Where("username = ?", username)
	if excludeDeleted {
		tx.Scopes(dao.NotDeleted)
	}
	var user User
	err := tx.First(&user).Error
	return &user, err
}

func (dao *Dao) LoadSystemUser() error {
	SystemUser = &User{
		Name:     "系统管理员",
		Username: viper.GetString("admin.user"),
		Status:   user_iface.UserStatus_USER_NORMAL,
	}
	SystemUser.SetPassword(viper.GetString("admin.pass"))
	return dao.Connection().Where("username = ?", SystemUser.Username).FirstOrCreate(SystemUser).Error
}
