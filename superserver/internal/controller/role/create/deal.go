package create

import (
	"go.uber.org/zap"
	"superserver/common/interface/ecode_iface"
	"superserver/common/interface/user_iface"
	"superserver/common/logger"
	"superserver/internal/model/role"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(Params)
	if params.Name == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "名称不能为空")
		return
	}
	roleDao := role.NewDao(ctl.GetDatabase())
	rolePo, err := roleDao.FindByName(params.Name)
	if err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "角色名称已存在")
		return
	}
	rolePo.Name = params.Name
	rolePo.Nature = user_iface.RoleNature_ROLE_CUSTOM
	err = roleDao.Create(&rolePo)
	if err != nil {
		logger.Error(ctl.Context, "create user roleDao.Create failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{})
}
