package utils

import (
	"fmt"
	"math/big"
	"net"
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
