package miniapp

import (
	"fmt"
	"net/http"
)

func (ctl *Controller) Deal() {
	fmt.Println(ctl.param.Command)
	ctl.NewOkResponse(http.StatusOK, Reply{
		Status:  "success",
		Cost:    "0.1ms",
		Message: fmt.Sprintf("(2006-01-02 15:03:04)[INFO] run command %s success", ctl.param.Command),
	})
}
