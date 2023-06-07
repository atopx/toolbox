package user_token

import (
	"cloudos/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (d *Dao) UpsertByUserId(token *UserToken) error {
	return d.Connection().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		UpdateAll: true,
	}).Create(&token).Error
}
