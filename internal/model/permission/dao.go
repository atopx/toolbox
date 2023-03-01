package permission

import (
	"superserver/internal/model"

	"gorm.io/gorm"
)

type PermissionDao struct {
	model.BaseDao
}

func (*PermissionDao) TableName() string {
	return "su_permission"
}

func NewDao(db *gorm.DB) *PermissionDao {
	dao := &PermissionDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *PermissionDao) Inspector(roleId, accessId int) bool {
	return dao.Connection().Where("role_id=? and access_id=?", roleId, accessId).
		First(&Permission{}).Error == nil
}
