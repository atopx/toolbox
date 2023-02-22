package delete

import (
	"errors"
	"net/http"
	"superserver/internal/model/computer"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	// TODO 验证是否有删除权限， 创建人或管理员

	dao := computer.NewDao(ctl.GetDatabase())
	po := computer.NewComputer(ctl.param.Id)
	err := dao.Load(po)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的主机ID")
		return
	}
	if err := dao.Delete(po); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "删除主机异常，请联系管理员")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
