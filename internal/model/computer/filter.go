package computer

import (
	"superserver/proto/common_iface"
	"superserver/proto/computer_iface"
)

type Filter struct {
	Keyword         string                               `json:"keyword"`
	Hostname        string                               `json:"hostname"`
	Address         string                               `json:"address"`
	PowerStatus     []computer_iface.ComputerPowerStatus `json:"power_status"`
	CreateTimeRange *common_iface.RangeI64               `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64               `json:"update_time_range"`
	ScanTimeRange   *common_iface.RangeI64               `json:"scan_time_range"`
}
