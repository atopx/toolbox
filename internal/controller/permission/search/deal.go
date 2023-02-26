package search

import (
	"go.uber.org/zap"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/access"
	"superserver/internal/model/permission"
	"superserver/internal/model/role"
	"superserver/internal/model/user"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	dao := permission.NewDao(ctl.GetDatabase())
	data, err := dao.Filter(&params.Filter, params.Page)
	if err != nil {
		logger.Error(ctl.Context, "search permission dao.Filter failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	var userIds, roleIds, accessIds []int
	for _, po := range data {
		userIds = append(userIds, po.Creator, po.Updater)
		roleIds = append(roleIds, po.RoleId)
		accessIds = append(accessIds, po.AccessId)
	}
	userMap, err := user.NewDao(ctl.GetDatabase()).GetUserMapByIds(userIds)
	if err != nil {
		logger.Error(ctl.Context, "search permission GetUserMapByIds failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	accessMap, err := access.NewDao(ctl.GetDatabase()).GetAccessMapByIds(accessIds)
	if err != nil {
		logger.Error(ctl.Context, "search permission GetAccessMapByIds failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	roleMap, err := role.NewDao(ctl.GetDatabase()).GetRoleMapByIds(roleIds)
	if err != nil {
		logger.Error(ctl.Context, "search permission GetRoleMapByIds failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}

	reply := Reply{
		Page: params.Page,
		List: []PermissionVo{},
	}
	for _, po := range data {
		vo := ctl.NewPermissionVo(&po)
		if value, ok := accessMap[po.AccessId]; ok {
			vo.Access = &AccessVo{
				Id:      value.Id,
				Path:    value.Path,
				Method:  value.Method,
				Handler: value.Handler,
			}
		}
		if value, ok := roleMap[po.RoleId]; ok {
			vo.Role = &RoleVo{
				Id:     value.Id,
				Name:   value.Name,
				Nature: value.Nature,
			}
		}
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

func (ctl *Controller) NewPermissionVo(po *permission.Permission) PermissionVo {
	return PermissionVo{
		Id:         po.Id,
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}
