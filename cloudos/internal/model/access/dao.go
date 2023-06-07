package access

import (
	"cloudos/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Dao struct {
	model.BaseDao
}

func (*Dao) TableName() string {
	return "su_access"
}

func NewDao(db *gorm.DB) *Dao {
	dao := &Dao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *Dao) Get(method, path string) (*Access, error) {
	access := new(Access)
	err := dao.Connection().Where("method = ? and path = ?", method, path).First(access).Error
	return access, err
}

func (dao *Dao) BatchUpsert(accessList []Access) error {
	return dao.Connection().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		DoUpdates: clause.AssignmentColumns([]string{"method", "handler", "update_time", "delete_time"}),
	}).Create(&accessList).Error
}

func (dao *Dao) GetAccessMapByIds(ids []int) (map[int]*Access, error) {
	result := make(map[int]*Access)
	list, err := dao.GetAccessListByIds(ids)
	if err != nil {
		return result, err
	}
	for _, access := range list {
		result[access.Id] = access
	}
	return result, nil
}

func (dao *Dao) GetAccessListByIds(ids []int) (list []*Access, err error) {
	if len(ids) == 0 {
		return list, nil
	}
	err = dao.Connection().Where("id in ?", ids).Find(&list).Error
	return list, err
}
