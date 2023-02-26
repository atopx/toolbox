package role

import (
	"errors"
	"superserver/common/utils"
	"superserver/internal/model"
	"superserver/proto/common_iface"
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

func (dao *RoleDao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Role, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().Where("delete_time = 0")
	if filter.Keyword != "" {
		tx.Where("name like ?", utils.NewLikeValue(filter.Keyword))
	}

	if len(filter.Nature) > 0 {
		tx.Where("nature in ?", filter.Nature)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&computers).Error
	return computers, err
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
