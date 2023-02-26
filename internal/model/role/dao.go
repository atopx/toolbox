package role

import (
	"errors"
	"superserver/internal/model"
	"superserver/proto/user_iface"

	"gorm.io/gorm"
)

type RoleDao struct {
	model.BaseDao
}

func (*RoleDao) TableName() string {
	return "su_role"
}

func NewDao(db *gorm.DB) *RoleDao {
	dao := &RoleDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *RoleDao) LoadSystemRole() (err error) {
	SystemRole, err = dao.loadByNature("系统管理员", user_iface.RoleNature_ROLE_SYSTEM)
	return err
}

func (dao *RoleDao) LoadDefaultRole() (err error) {
	DefaultRole, err = dao.loadByNature("默认角色", user_iface.RoleNature_ROLE_DEFAULT)
	return err
}

func (dao *RoleDao) loadByNature(name string, nature user_iface.RoleNature) (*Role, error) {
	role := new(Role)
	err := dao.Connection().Where("nature=?", nature).First(role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		role.Name = "系统管理员"
		role.Nature = user_iface.RoleNature_ROLE_DEFAULT
		err = dao.Create(role)
	}
	return role, err
}
