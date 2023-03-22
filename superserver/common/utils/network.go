package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh"
	"math/big"
	"net"
	"strings"
	"superserver/common/interface/computer_iface"
	"time"
)

func IPDecode(value int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func IPEncode(ip string) int64 {
	value := big.NewInt(0)
	value.SetBytes(net.ParseIP(ip).To4())
	return value.Int64()
}

var macReplacer = strings.NewReplacer("-", "", ":", "")

func MACEncode(mac string) string {
	return macReplacer.Replace(mac)
}

func MACDecode(mac string) string {
	var values []string
	for i := 0; i < len(mac); i += 2 {
		values = append(values, mac[i:i+2])
	}
	return strings.Join(values, ":")
}

func ComputerCheck(host string, port int) computer_iface.ComputerPowerStatus {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Second*10)
	if err != nil {
		return computer_iface.ComputerPowerStatus_COMPUTER_POWER_OFF
	}
	_ = conn.Close()
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
