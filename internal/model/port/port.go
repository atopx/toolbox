package port

import (
	"superserver/internal/model"
	"superserver/proto/computer_iface"
	"sync"
)

type ComputerPort struct {
	model.Base
	ComputerId int                                  // 主机ID
	Port       uint16                               // 网络端口
	Protocol   computer_iface.ComputerPortProtocol  // 应用协议
	Transport  computer_iface.ComputerPortTransport // 传输协议
	UseType    computer_iface.ComputerPortUseType   // 用途
	Desc       string                               // 备注
	Creator    int                                  // 创建人
	Updator    int                                  // 更新人
}

var once sync.Once
