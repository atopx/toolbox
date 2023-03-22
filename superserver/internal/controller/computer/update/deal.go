package update

import (
	"errors"
	"superserver/common/interface/ecode_iface"
	"superserver/internal/model/computer"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	// TODO 验证是否有删除权限， 创建人或管理员

	dao := computer.NewDao(ctl.GetDatabase())
	po := computer.NewComputer(params.Id)
	err := dao.Load(po)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "无效的主机ID")
		return
	}
	po.Name = params.Name
	po.Username = params.Name
	po.Password = params.Name
	po.LanHostname = params.Name
	po.WanHostname = params.Name
	po.Address = params.Address
	if err := dao.Update(po); err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "更新主机异常，请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{})
}
