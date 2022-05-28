package spec

import (
	"math"
	"net/http"
	"toolbox/internal/model/pixel"
)

func (ctl *Controller) Deal() {
	ppi := math.Sqrt(math.Pow(ctl.param.WidthRate, 2)+math.Pow(ctl.param.HeightRate, 2)) / ctl.param.Size
	reply := &Reply{
		Size:   ctl.param.Size,
		PPI:    ppi,
		Width:  pixel.Spec{Rate: ctl.param.WidthRate},
		Height: pixel.Spec{Rate: ctl.param.HeightRate},
	}
	reply.Width.SetLength(ppi)
	reply.Height.SetLength(ppi)
	ctl.NewOkResponse(http.StatusOK, reply)
}
