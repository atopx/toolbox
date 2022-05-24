package wechat

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

const token = "atopx"

type IntoParam struct {
	Signature string `json:"Signature" form:"signature"`
	Timestamp string `json:"Timestamp" form:"timestamp"`
	Nonce     string `json:"Nonce" form:"nonce"`
	Echostr   string `json:"Echostr" form:"echostr"`
}

func (param *IntoParam) Validator() bool {
	tmp := []string{token, param.Timestamp, param.Nonce}
	sort.Slice(tmp, func(i, j int) bool {
		return i > j
	})
	hash := sha1.New()
	hash.Write([]byte(strings.Join(tmp, "")))
	return param.Signature == fmt.Sprintf("%x", hash.Sum(nil))
}
