package role

import (
	"superserver/proto/common_iface"
	"superserver/proto/user_iface"
)

type Filter struct {
	Ids             []int                   `json:"ids"`
	Keyword         string                  `json:"keyword"`
	Nature          []user_iface.RoleNature `json:"nature"`
	CreateTimeRange *common_iface.RangeI64  `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64  `json:"update_time_range"`
}
