package computer

import (
	"superserver/common/interface/computer_iface"
	"superserver/common/utils"
	"superserver/internal/model"
)

type Computer struct {
	model.Base

	Name        string                             // 名称
	Username    string                             // 用户名
	Password    string                             // 用户密码
	LanHostname string                             // 局域网地址
	WanHostname string                             // 广域网地址
	Address     string                             // 物理地址
	PowerStatus computer_iface.ComputerPowerStatus // 电源状态
	Creator     int                                // 创建人
	Updater     int                                // 更新人
	ScanTime    int64                              // 最后一次扫描时间
}

func NewComputer(id int) *Computer {
	return &Computer{Base: model.Base{Id: id}}
}

func (m *Computer) SetAddress(address string) {
	m.Address = utils.MACEncode(address)
}

func (m *Computer) GetAddress() string {
	return utils.MACDecode(m.Address)
}
