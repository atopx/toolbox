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

func (dao *RoleDao) FindByName(name string) (role Role, err error) {
	err = dao.Connection().Where("name = ?", name).First(&role).Error
	return role, err
}

func (dao *RoleDao) GetRoleMapByIds(ids []int) (map[int]*Role, error) {
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
		role.Name = name
		role.Nature = nature
		err = dao.Create(role)
	}
	return role, err
}
