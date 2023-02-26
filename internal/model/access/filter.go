package access

import (
	"superserver/common/utils"
	"superserver/proto/common_iface"
)

type Filter struct {
	Ids             []int                       `json:"ids"`
	Keyword         string                      `json:"keyword"`
	Method          []common_iface.AccessMethod `json:"method"`
	CreateTimeRange *common_iface.RangeI64      `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64      `json:"update_time_range"`
}

func (dao *AccessDao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Access, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().Where("delete_time = 0")
	if filter.Keyword != "" {
		tx.Where("path like ? or handler like ?", utils.NewLikeValue(filter.Keyword), utils.NewLikeValue(filter.Keyword))
	}

	if len(filter.Method) > 0 {
		tx.Where("method in ?", filter.Method)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&computers).Error
	return computers, err
}
