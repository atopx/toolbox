package search

import (
	"go.uber.org/zap"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/computer"
	"superserver/internal/model/user"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	dao := computer.NewDao(ctl.GetDatabase())
	data, err := dao.Filter(&params.Filter, params.Page)
	if err != nil {
		logger.Error(ctl.Context, "search computer dao.Filter failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}

	var userIds []int
	for _, po := range data {
		userIds = append(userIds, po.Creator, po.Updater)
	}
	userMap, err := user.NewDao(ctl.GetDatabase()).GetUserMapByIds(userIds)
	if err != nil {
		logger.Error(ctl.Context, "search computer GetUserMapByIds failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	reply := &Reply{
		Page: params.Page,
		List: []ComputerVo{},
	}

	for _, po := range data {
		vo := ctl.NewComputerVo(&po)
		if value, ok := userMap[po.Creator]; ok {
			vo.Creator = value.Name
		}
		if value, ok := userMap[po.Updater]; ok {
			vo.Updater = value.Name
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
		Address:     po.GetAddress(),
		PowerStatus: po.PowerStatus.String(),
		CreateTime:  po.CreateTime,
		UpdateTime:  po.UpdateTime,
		ScanTime:    po.ScanTime,
	}
}
