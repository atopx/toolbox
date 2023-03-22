package update

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"superserver/common/interface/ecode_iface"
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
	_, err := roleDao.FindByName(params.Name)
	if err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "角色名称已存在")
		return
	}
	rolePo := role.New(params.Id)
	if errors.Is(roleDao.Load(rolePo), gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "角色不存在")
		return
	}
	rolePo.Name = params.Name
	err = roleDao.Update(rolePo)
	if err != nil {
		logger.Error(ctl.Context, "update role roleDao.Update failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{})
}
