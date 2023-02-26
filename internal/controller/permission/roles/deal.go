package roles

import (
	"net/http"
)

func (ctl *Controller) Deal() {
	reply := &Reply{}
	ctl.NewOkResponse(http.StatusOK, &reply)
}
