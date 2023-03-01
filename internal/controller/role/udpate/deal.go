package update

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
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
	_, err := roleDao.FindByName(params.Name)
	if err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, "角色名称已存在")
		return
	}
	rolePo := role.New(params.Id)
	if errors.Is(roleDao.Load(rolePo), gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "角色不存在")
		return
	}
	rolePo.Name = params.Name
	err = roleDao.Update(rolePo)
	if err != nil {
		logger.Error(ctl.Context, "update role roleDao.Update failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
