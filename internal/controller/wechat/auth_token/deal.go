package auth_token

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

func (ctl *Controller) Deal() {
	tmp := []string{viper.GetString("wechat.token"), ctl.param.Timestamp, ctl.param.Nonce}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	hash := sha1.New()
	hash.Write([]byte(strings.Join(tmp, "")))
	signature := hex.EncodeToString(hash.Sum(nil))
	if ctl.param.Signature != signature {
		ctl.NewErrorResponse(http.StatusForbidden, fmt.Sprintf("expect: %s, current: %s", ctl.param.Signature, signature))
		return
	}
	ctl.Context.String(http.StatusOK, ctl.param.Echostr)
}
