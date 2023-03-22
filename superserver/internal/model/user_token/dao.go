package user_token

import (
	"gorm.io/gorm"
	"superserver/internal/model"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_user_token"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (d *Dao) First(where func(*gorm.DB) *gorm.DB) (*UserToken, error) {
	token := new(UserToken)
	err := d.Connection().Scopes(where).First(&token).Error
	return token, err
}
