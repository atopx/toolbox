package user

import (
	"gorm.io/gorm"
	"superserver/internal/model"
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

func (dao *UserDao) Upsert(user *User) error {
	return dao.Connection().Where(User{Username: SystemUser.Username, Role: SystemUser.Role}).FirstOrCreate(user).Error
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
