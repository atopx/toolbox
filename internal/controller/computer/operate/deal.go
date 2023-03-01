package operate

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"net/http"
	"superserver/internal/model/computer"
	"superserver/proto/computer_iface"
	"time"

	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Operate == nil || params.Id <= 0 {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的操作")
		return
	}

	// TODO 验证是否有操作权限， 创建人或管理员

	// 验证主机是否有效
	dao := computer.NewDao(ctl.GetDatabase())
	po := computer.NewComputer(params.Id)
	err := dao.Load(po)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的主机ID")
		return
	}

	switch *params.Operate {
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_OFF:
		// TODO 关机
		err = ComputerOff(po.Username, po.Password, po.LanHostname, 22)
		fmt.Println(computer_iface.ComputerOperate_COMPUTER_OPERATE_OFF.String(), err)
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_ON:
		err = ComputerOn(po.Address)
		fmt.Println(computer_iface.ComputerOperate_COMPUTER_OPERATE_ON.String(), err)
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_DETECT:
		status := ComputerCheck("tcp", po.LanHostname, 3389)
		fmt.Println(computer_iface.ComputerOperate_COMPUTER_OPERATE_DETECT.String(), status)
	}
	if err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, "操作失败, 请检查配置")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}

func ComputerCheck(transport, host string, port int) computer_iface.ComputerPowerStatus {
	conn, err := net.DialTimeout(transport, fmt.Sprintf("%s:%d", host, port), time.Second*10)
	if err != nil {
		return computer_iface.ComputerPowerStatus_COMPUTER_POWER_OFF
	}
	conn.Close()
	return computer_iface.ComputerPowerStatus_COMPUTER_POWER_ON
}

func ComputerOn(mac string) error {
	machex, _ := hex.DecodeString(mac)
	conn, err := net.DialUDP("udp", &net.UDPAddr{}, &net.UDPAddr{IP: net.IPv4bcast})
	if err != nil {
		return fmt.Errorf("network error: %s", err)
	}
	defer conn.Close()
	var buffer bytes.Buffer
	buffer.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	for i := 0; i < 16; i++ {
		buffer.Write(machex)
	}
	if _, err = conn.Write(buffer.Bytes()); err != nil {
		return fmt.Errorf("send packet error: %s", err)
	}
	return nil
}

func ComputerOff(user, pass, host string, port int) error {
	config := &ssh.ClientConfig{
		Timeout:         10 * time.Second,
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return fmt.Errorf("connect server failed: %s", err.Error())
	}
	defer client.Close()
	sess, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("login shell failed: %s", err.Error())
	}
	defer sess.Close()
	return sess.Start("shutdown -s -t 0")
}
