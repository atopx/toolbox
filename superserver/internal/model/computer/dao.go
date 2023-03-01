package computer

import (
	"fmt"
	"gorm.io/gorm"
	"superserver/internal/model"
)

type ComputerDao struct {
	model.BaseDao
}

func (*ComputerDao) TableName() string {
	return "su_computer"
}

func NewDao(db *gorm.DB) *ComputerDao {
	dao := &ComputerDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *ComputerDao) FilterBy(key string, value any) (po *Computer, err error) {
	err = dao.Connection().Where(fmt.Sprintf("%s = ?", key), value).First(&po).Error
	return po, err
}
