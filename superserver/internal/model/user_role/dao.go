package user_role

import (
	"gorm.io/gorm"
	"superserver/internal/model"
)

type UserRoleRefDao struct {
	model.BaseDao
}

func (*UserRoleRefDao) TableName() string {
	return "su_user_role_ref"
}

func NewDao(db *gorm.DB) *UserRoleRefDao {
	dao := &UserRoleRefDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}
