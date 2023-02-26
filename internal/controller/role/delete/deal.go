package delete

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/role"
	"superserver/proto/user_iface"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(Params)
	if params.Id == 0 {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的角色")
		return
	}
	roleDao := role.NewDao(ctl.GetDatabase())
	rolePo := role.New(params.Id)
	err := roleDao.Load(rolePo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "角色不存在")
		return
	}
	if err != nil {
		logger.Error(ctl.Context, "delete role roleDao.Load failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
		return
	}
	if rolePo.Nature != user_iface.RoleNature_ROLE_CUSTOM {
		ctl.NewErrorResponse(http.StatusBadRequest, "内置角色禁止删除")
		return
	}
	if err = roleDao.Delete(rolePo); err != nil {
		logger.Error(ctl.Context, "delete role roleDao.Delete failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
