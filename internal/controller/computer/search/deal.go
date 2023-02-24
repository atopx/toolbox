package search

import (
	"net/http"
	"superserver/internal/model/computer"
	"superserver/internal/model/user"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	dao := computer.NewDao(ctl.GetDatabase())
	data, err := dao.Filter(&params.Filter, params.Page)
	if err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "[10000]系统错误,请联系管理员")
		return
	}

	var userIds []int
	for _, po := range data {
		userIds = append(userIds, po.Creator, po.Updator)
	}
	userMap, err := user.NewDao(ctl.GetDatabase()).GetUserMapByIds(userIds)
	if err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "[10001]系统错误,请联系管理员")
		return
	}
	reply := &Reply{
		Page: params.Page,
		List: []ComputerVo{},
	}

	for _, po := range data {
		vo := ctl.NewComputerVo(&po)
		if user, ok := userMap[po.Creator]; ok {
			vo.Creator = user.Name
		}
		if user, ok := userMap[po.Updator]; ok {
			vo.Updator = user.Name
		}
		reply.List = append(reply.List, vo)
	}
	ctl.NewOkResponse(http.StatusOK, reply)
}

func (ctl *Controller) NewComputerVo(po *computer.Computer) ComputerVo {
	return ComputerVo{
		Id:          po.Id,
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
	}
}
