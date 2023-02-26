package access

import (
	"fmt"
	"net/http"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	fmt.Println(params)
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
