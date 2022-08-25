package compress

import (
	"net/http"
	"strings"
)

var replacer = strings.NewReplacer(" ", "")

func (ctl *Controller) Deal() {
	ctl.Context.String(http.StatusOK, replacer.Replace(ctl.param.JsonStr))
}
