package port

import (
	"cloudos/internal/model"

	"gorm.io/gorm"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_computer_port"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}
