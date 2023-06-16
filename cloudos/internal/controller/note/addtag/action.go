package addtag

import (
	"cloudos/common/consts"
	"cloudos/common/pb"
	"fmt"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)
	fmt.Println(params)

	return consts.EmptyObj, pb.ECode_SUCCESS
}
