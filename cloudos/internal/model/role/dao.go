package role

import (
	"cloudos/common/interface/user_iface"
	"cloudos/internal/model"

	"gorm.io/gorm"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_role"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *Dao) FindByName(name string) (role Role, err error) {
	err = dao.Connection().Where("name = ?", name).First(&role).Error
	return role, err
}

func (dao *Dao) GetRoleMapByIds(ids []int) (map[int]*Role, error) {
	result := make(map[int]*Role)
	if len(ids) == 0 {
		return result, nil
	}
	var roles []*Role
	err := dao.Connection().Where("id in ?", ids).Find(&roles).Error
	if err != nil {
		return result, err
	}
	for _, role := range roles {
		result[role.Id] = role
	}
	return result, nil
}

func (dao *Dao) LoadSystemRole() (err error) {
	SystemRole, err = dao.loadByNature("超级管理员", user_iface.RoleNature_ROLE_SYSTEM)
	return err
}

func (dao *Dao) LoadDefaultRole() (err error) {
	DefaultRole, err = dao.loadByNature("默认角色", user_iface.RoleNature_ROLE_DEFAULT)
	return err
}

func (dao *Dao) loadByNature(name string, nature user_iface.RoleNature) (*Role, error) {
	role := &Role{Name: name, Nature: nature}
	err := dao.Connection().Where("nature=?", nature).FirstOrCreate(role).Error
	return role, err
}
