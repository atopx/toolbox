package wol

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
)

var macAddrMatch = regexp.MustCompile(`^([0-9a-fA-F]{2}[:-]){5}[0-9a-fA-F]{2}$`)
var macAddrReplacer = strings.NewReplacer("-", "", ":", "")

func (ctl *Controller) Deal() {

	// 参数校验
	if !macAddrMatch.MatchString(ctl.param.MacAddr) {
		ctl.NewErrorResponse(http.StatusBadRequest, "Invalid mac address")
		return
	}
	if err := ctl.packeter(); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, err.Error())
		return
	}
	ctl.NewOkResponse(http.StatusOK, Reply{"success"})
}

func (ctl *Controller) packeter() error {
	machex, _ := hex.DecodeString(macAddrReplacer.Replace(ctl.param.MacAddr))
	conn, err := net.DialUDP("udp", &net.UDPAddr{}, &net.UDPAddr{IP: net.IPv4bcast})
	if err != nil {
		// ctl.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("network error: %s", err))
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
