package user

import (
	"superserver/internal/model"

	"gorm.io/gorm"
)

type UserDao struct {
	model.BaseDao
}

func NewDao(db *gorm.DB) *UserDao {
	dao := &UserDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.Tablename = dao.Tablename()
	return dao
}

func (*UserDao) Tablename() string {
	return "supan_user"
}

func (dao *UserDao) Create(user *User) error {
	return dao.BaseDao.Create(user)
}

func (dao *UserDao) Login(user *User) error {
	return dao.Connection().Where("user=? and password=?", user.Username, user.Password).First(user).Error
}
