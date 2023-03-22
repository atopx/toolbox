package delete

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
	if err := dao.Delete(po); err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "删除主机异常，请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{})
}
