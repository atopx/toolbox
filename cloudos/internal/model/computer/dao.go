package computer

import (
	"cloudos/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_computer"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *Dao) FilterBy(key string, value any) (po *Computer, err error) {
	err = dao.Connection().Where(fmt.Sprintf("%s = ?", key), value).First(&po).Error
	return po, err
}
