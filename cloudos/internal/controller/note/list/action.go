package list

import (
	"cloudos/common/pb"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)

	reply := Reply{
		Pager: params.Pager,
		List:  make([]Item, 0),
	}

	return reply, pb.ECode_SUCCESS
}
