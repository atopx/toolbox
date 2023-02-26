package port

import (
	"superserver/internal/model"

	"gorm.io/gorm"
)

type PortDao struct {
	model.BaseDao
}

func (*PortDao) TableName() string {
	return "su_computer_port"
}

func NewDao(db *gorm.DB) *PortDao {
	dao := &PortDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}
