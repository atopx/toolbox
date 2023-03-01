package create

import (
	"go.uber.org/zap"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/permission"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.RoleId == 0 || params.AccessId == 0 {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的操作")
		return
	}
	permissionDao := permission.NewDao(ctl.GetDatabase())
	if permissionDao.Inspector(params.RoleId, params.AccessId) {
		ctl.NewErrorResponse(http.StatusBadRequest, "角色权限已存在")
		return
	}
	err := permissionDao.Create(&permission.Permission{
		AccessId: params.AccessId,
		RoleId:   params.RoleId,
		Creator:  ctl.GetOperator(),
		Updater:  ctl.GetOperator(),
	})
	if err != nil {
		logger.Error(ctl.Context, "create permission permissionDao.Create failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusBadRequest, "角色名称已存在")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
