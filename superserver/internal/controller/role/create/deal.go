package create

import (
	"go.uber.org/zap"
	"net/http"
	"superserver/common/interface/user_iface"
	"superserver/common/logger"
	"superserver/internal/model/role"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(Params)
	if params.Name == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "名称不能为空")
		return
	}
	roleDao := role.NewDao(ctl.GetDatabase())
	rolePo, err := roleDao.FindByName(params.Name)
	if err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, "角色名称已存在")
		return
	}
	rolePo.Name = params.Name
	rolePo.Nature = user_iface.RoleNature_ROLE_CUSTOM
	err = roleDao.Create(&rolePo)
	if err != nil {
		logger.Error(ctl.Context, "create user roleDao.Create failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
