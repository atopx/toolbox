package update

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
	po.Name = ctl.param.Name
	po.Username = ctl.param.Name
	po.Password = ctl.param.Name
	po.LanHostname = ctl.param.Name
	po.WanHostname = ctl.param.Name
	po.Address = ctl.param.Address
	if err := dao.Update(po); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "更新主机异常，请联系管理员")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
