package permission

import (
	"cloudos/internal/model"
	"errors"

	"gorm.io/gorm"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_permission"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *Dao) Inspector(userId, accessId int) bool {
	return errors.Is(dao.Connection().Select("id").
		Where("exists(select id from su_user_role_ref where user_id = ? and role_id = su_permission.role_id)", userId).
		Where("access_id = ?", accessId).First(&Permission{}).Error, gorm.ErrRecordNotFound)
}
