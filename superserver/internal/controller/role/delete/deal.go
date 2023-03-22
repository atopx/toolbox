package delete

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"superserver/common/interface/ecode_iface"
	"superserver/common/interface/user_iface"
	"superserver/common/logger"
	"superserver/internal/model/role"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(Params)
	if params.Id == 0 {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "无效的角色")
		return
	}
	roleDao := role.NewDao(ctl.GetDatabase())
	rolePo := role.New(params.Id)
	err := roleDao.Load(rolePo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "角色不存在")
		return
	}
	if err != nil {
		logger.Error(ctl.Context, "delete role roleDao.Load failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}
	if rolePo.Nature != user_iface.RoleNature_ROLE_CUSTOM {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "内置角色禁止删除")
		return
	}
	if err = roleDao.Delete(rolePo); err != nil {
		logger.Error(ctl.Context, "delete role roleDao.Delete failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{})
}
