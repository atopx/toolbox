package computer

import (
	"fmt"
	"superserver/common/utils"
	"superserver/internal/model"
	"superserver/proto/common_iface"

	"gorm.io/gorm"
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

func (dao *ComputerDao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Computer, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().Where("delete_time = 0")
	if filter.Keyword != "" {
		tx.Where("name like ? or username like ?", utils.NewLikeValue(filter.Keyword), utils.NewLikeValue(filter.Keyword))
	}

	if filter.Hostname != "" {
		tx.Where("lan_hostname like ? or wan_hostname like ?", utils.NewLikeValue(filter.Hostname), utils.NewLikeValue(filter.Hostname))
	}

	if filter.Address != "" {
		filter.Address = utils.MACEncode(filter.Address)
		tx.Where("address like ?", utils.NewLikeValue(filter.Address))
	}

	if len(filter.PowerStatus) > 0 {
		tx.Where("power_status in ?", filter.PowerStatus)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Range("scan_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&computers).Error
	return computers, err
}
