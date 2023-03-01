package role

import (
	"superserver/common/utils"
	"superserver/proto/common_iface"
	"superserver/proto/user_iface"
)

type Filter struct {
	Ids             []int                   `json:"ids"`
	Keyword         string                  `json:"keyword"`
	Natures         []user_iface.RoleNature `json:"natures"`
	CreateTimeRange *common_iface.RangeI64  `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64  `json:"update_time_range"`
}

func (dao *RoleDao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Role, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().Scopes(dao.NotDeleted)
	if filter.Keyword != "" {
		tx.Where("name like ?", utils.NewLikeValue(filter.Keyword))
	}
	if len(filter.Ids) > 0 {
		tx.Where("id in ?", filter.Ids)
	}
	if len(filter.Natures) > 0 {
		tx.Where("nature in ?", filter.Natures)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&computers).Error
	return computers, err
}
