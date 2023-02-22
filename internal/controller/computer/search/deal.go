package search

import (
	"net/http"
	"superserver/internal/model/computer"
)

func (ctl *Controller) Deal() {
	dao := computer.NewDao(ctl.GetDatabase())
	data, err := dao.Filter(&ctl.param.Filter, ctl.param.Page)
	if err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	reply := &Reply{
		Page: ctl.param.Page,
		List: make([]ComputerVo, len(data)),
	}
	for i, po := range data {
		reply.List[i] = ctl.NewComputerVo(&po)
	}
	ctl.NewOkResponse(http.StatusOK, reply)
}

func (ctl *Controller) NewComputerVo(po *computer.Computer) ComputerVo {
	vo := ComputerVo{
		Name:        po.Name,
		Username:    po.Username,
		Password:    po.Password,
		LanHostname: po.LanHostname,
		WanHostname: po.WanHostname,
		Address:     po.Address,
		PowerStatus: po.PowerStatus.String(),
		CreateTime:  po.CreateTime,
		UpdateTime:  po.UpdateTime,
		ScanTime:    po.ScanTime,

		// TODO
		// Creator:     po.Creator,
		// Updator:     po.Updator,
	}
	return vo
}
