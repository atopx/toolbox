package permission

import (
	"superserver/common/utils"
	"superserver/proto/common_iface"
	"superserver/proto/user_iface"
)

type Filter struct {
	Ids             []int                       `json:"ids"`
	RoleIds         []int                       `json:"role_ids"`
	AccessIds       []int                       `json:"access_ids"`
	Keyword         string                      `json:"keyword"`
	AccessMethods   []common_iface.AccessMethod `json:"access_methods"`
	RoleNatures     []user_iface.RoleNature     `json:"role_natures"`
	CreateTimeRange *common_iface.RangeI64      `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64      `json:"update_time_range"`
}

func (dao *PermissionDao) Filter(filter *Filter, pager *common_iface.Pager) (permissions []Permission, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().
		Joins("join role on permission.role_id = role.id and role.delete_time > 0").
		Joins("join access on permission.access_id = access.id and access.delete_time > 0")

	if filter.Keyword != "" {
		keyword := utils.NewLikeValue(filter.Keyword)
		tx.Where("access.path like ? or access.handler like ? or role.name like ?", keyword, keyword, keyword)
	}
	if len(filter.Ids) > 0 {
		tx.Where("permission.id in ?", filter.RoleIds)
	}

	if len(filter.RoleIds) > 0 {
		tx.Where("role.id in ?", filter.RoleIds)
	}
	if len(filter.RoleNatures) > 0 {
		tx.Where("role.nature in ?", filter.RoleNatures)
	}

	if len(filter.AccessIds) > 0 {
		tx.Where("access.id in ?", filter.RoleIds)
	}
	if len(filter.AccessMethods) > 0 {
		tx.Where("access.method in ?", filter.AccessMethods)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&permissions).Error
	return permissions, err
}
