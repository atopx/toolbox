package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func ComputerOn() error {
	machex, _ := hex.DecodeString("D8BBC1DFBEFC")
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

func ComputerOff() error {
	config := &ssh.ClientConfig{
		Timeout:         10 * time.Second,
		User:            "atopx",
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{ssh.Password("123456")},
	}

	client, err := ssh.Dial("tcp", "192.168.0.100:22", config)
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

type ComputerStatus string

const (
	ComputerStatusOnline  ComputerStatus = "ONLINE"
	ComputerStatusOffline ComputerStatus = "OFFLINE"
)

func ComputerCheck() ComputerStatus {
	conn, err := net.DialTimeout("tcp", "home.itmeng.top:3389", time.Second*10)
	if err != nil {
		return ComputerStatusOffline
	}
	conn.Close()
	return ComputerStatusOnline
}

func main() {
	fmt.Println("current computer status:", ComputerCheck())
}
