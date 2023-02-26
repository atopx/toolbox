package access

import (
	"superserver/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccessDao struct {
	model.BaseDao
}

func (*AccessDao) TableName() string {
	return "su_access"
}

func NewDao(db *gorm.DB) *AccessDao {
	dao := &AccessDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.TableName = dao.TableName()
	return dao
}

func (dao *AccessDao) BatchUpsert(accessList []Access) error {
	return dao.Connection().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		DoUpdates: clause.AssignmentColumns([]string{"method", "handler", "update_time", "delete_time"}),
	}).Create(&accessList).Error
}

func (dao *AccessDao) GetAccessMapByIds(ids []int) (map[int]*Access, error) {
	result := make(map[int]*Access)
	if len(ids) == 0 {
		return result, nil
	}
	var accessList []*Access
	err := dao.Connection().Where("id in ?", ids).Find(&accessList).Error
	if err != nil {
		return result, err
	}
	for _, access := range accessList {
		result[access.Id] = access
	}
	return result, nil
}
