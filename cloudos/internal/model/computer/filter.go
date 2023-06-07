package computer

import (
	"cloudos/common/interface/common_iface"
	"cloudos/common/interface/computer_iface"
	"cloudos/common/utils"
)

type Filter struct {
	Ids             []int                                `json:"ids"`
	Keyword         string                               `json:"keyword"`
	Hostname        string                               `json:"hostname"`
	Address         string                               `json:"address"`
	PowerStatus     []computer_iface.ComputerPowerStatus `json:"power_status"`
	CreateTimeRange *common_iface.RangeI64               `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64               `json:"update_time_range"`
	ScanTimeRange   *common_iface.RangeI64               `json:"scan_time_range"`
}

func (dao *Dao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Computer, err error) {
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
		keyword := utils.NewLikeValue(filter.Hostname)
		tx.Where("lan_hostname like ? or wan_hostname like ?", keyword, keyword)
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
