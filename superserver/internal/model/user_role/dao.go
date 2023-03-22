package user_role

import (
	"gorm.io/gorm"
	"superserver/internal/model"
	"superserver/internal/model/role"
	"superserver/internal/model/user"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_user_role_ref"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (d *Dao) GetUsers(roleId int) (users []user.User, err error) {
	sub := d.Connection().Select("user_id").Where("role_id = ?", roleId)
	err = user.NewDao(d.Db).Connection().Where("id in (?)", sub).Find(&users).Error
	return users, err
}

func (d *Dao) GetRoles(userId int) (roles []role.Role, err error) {
	sub := d.Connection().Select("role_id").Where("user_id = ?", userId)
	err = role.NewDao(d.Db).Connection().Where("id in (?)", sub).Find(&roles).Error
	return roles, err
}
