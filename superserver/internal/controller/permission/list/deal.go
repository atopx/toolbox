package list

import (
	"fmt"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	fmt.Println(params)
	ctl.NewOkResponse(&Reply{})
}
