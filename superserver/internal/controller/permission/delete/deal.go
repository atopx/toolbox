package delete

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"superserver/common/interface/ecode_iface"
	"superserver/common/logger"
	"superserver/internal/model/permission"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(Params)
	if params.Id == 0 {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "无效的角色")
		return
	}
	permissionDao := permission.NewDao(ctl.GetDatabase())
	permissionPo := permission.New(params.Id)
	err := permissionDao.Load(permissionPo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "权限配置不存在")

		return
	}
	if err != nil {
		logger.Error(ctl.Context, "delete permission permissionDao.Load failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}
	if err = permissionDao.RealDelete(permissionPo); err != nil {
		logger.Error(ctl.Context, "delete permission roleDao.Delete failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{})
}
