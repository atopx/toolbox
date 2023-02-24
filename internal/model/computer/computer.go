package computer

import (
	"superserver/internal/model"
	"superserver/proto/computer_iface"
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
	Updator     int                                // 更新人
	ScanTime    int64                              // 最后一次扫描时间
}

func (m *Computer) GetId() any {
	return m.Id
}

func NewComputer(id int) *Computer {
	return &Computer{Base: model.Base{Id: id}}
}
