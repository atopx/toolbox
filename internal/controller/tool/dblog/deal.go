package dblog

import (
	"github.com/atotto/clipboard"
	"net/http"
	"strings"
)

var senseReplacer = strings.NewReplacer("\n", " ", "\t", " ")

func (ctl *Controller) Deal() {
	var result string
	temp := strings.Split(ctl.param.Message, "] ")
	if len(temp) > 1 {
		result = strings.TrimSpace(temp[len(temp)-1])
		result = senseReplacer.Replace(result) + ";"
	}
	_ = clipboard.WriteAll(result)
	ctl.NewOkResponse(http.StatusOK, result)
}
