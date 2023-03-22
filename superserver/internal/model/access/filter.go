package access

import (
	"superserver/common/interface/common_iface"
	"superserver/common/utils"
)

type Filter struct {
	Ids             []int                       `json:"ids"`
	Keyword         string                      `json:"keyword"`
	Methods         []common_iface.AccessMethod `json:"methods"`
	CreateTimeRange *common_iface.RangeI64      `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64      `json:"update_time_range"`
}

func (dao *Dao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Access, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().Scopes(dao.NotDeleted)
	if filter.Keyword != "" {
		keyword := utils.NewLikeValue(filter.Keyword)
		tx.Where("path like ? or handler like ?", keyword, keyword)
	}
	if len(filter.Ids) > 0 {
		tx.Where("id in ?", filter.Ids)
	}
	if len(filter.Methods) > 0 {
		tx.Where("method in ?", filter.Methods)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&computers).Error
	return computers, err
}
