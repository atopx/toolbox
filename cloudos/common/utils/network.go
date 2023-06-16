package utils

import (
	"cloudos/common/consts"
	"fmt"
	"math/big"
	"net"
	"strings"
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

var macReplacer = strings.NewReplacer("-", consts.EmptyStr, ":", consts.EmptyStr)

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
