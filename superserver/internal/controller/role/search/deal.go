package search

import (
	"go.uber.org/zap"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/role"
	"superserver/internal/model/user"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	dao := role.NewDao(ctl.GetDatabase())
	data, err := dao.Filter(&params.Filter, params.Page)
	if err != nil {
		logger.Error(ctl.Context, "search role dao.Filter failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	var userIds []int
	for _, po := range data {
		userIds = append(userIds, po.Creator, po.Updater)
	}
	userMap, err := user.NewDao(ctl.GetDatabase()).GetUserMapByIds(userIds)
	if err != nil {
		logger.Error(ctl.Context, "search role GetUserMapByIds failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	reply := Reply{
		Page: params.Page,
		List: []RoleVo{},
	}
	for _, po := range data {
		vo := ctl.NewRoleVo(&po)
		if value, ok := userMap[po.Creator]; ok {
			vo.Creator = value.Name
		}
		if value, ok := userMap[po.Updater]; ok {
			vo.Updater = value.Name
		}
		reply.List = append(reply.List, vo)
	}

	ctl.NewOkResponse(http.StatusOK, &reply)
}

func (ctl *Controller) NewRoleVo(po *role.Role) RoleVo {
	return RoleVo{
		Id:         po.Id,
		Name:       po.Name,
		Nature:     po.Nature,
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}
