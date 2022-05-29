package auth_token

import (
	"crypto/sha1"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"sort"
	"strings"
)

func (ctl *Controller) Deal() {
	var reply string
	tmp := []string{viper.GetString("wechat.token"), ctl.param.Timestamp, ctl.param.Nonce}
	sort.Slice(tmp, func(i, j int) bool {
		return i > j
	})
	hash := sha1.New()
	hash.Write([]byte(strings.Join(tmp, "")))
	if ctl.param.Signature != fmt.Sprintf("%x", hash.Sum(nil)) {
		ctl.NewErrorResponse(http.StatusForbidden, reply)
		return
	}
	ctl.NewOkResponse(http.StatusOK, reply)
}
