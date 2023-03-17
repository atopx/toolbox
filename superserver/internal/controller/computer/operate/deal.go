package operate

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"superserver/common/utils"
	"superserver/internal/model/computer"
	"superserver/proto/computer_iface"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Operate == nil || params.Id <= 0 {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的操作")
		return
	}
	var reply Reply

	// TODO 验证是否有操作权限， 创建人或管理员

	// 验证主机是否有效
	dao := computer.NewDao(ctl.GetDatabase())
	po := computer.NewComputer(params.Id)
	err := dao.Load(po)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的主机ID")
		return
	}

	switch *params.Operate {
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_OFF:
		err = utils.ComputerOff(po.Username, po.Password, po.LanHostname, 22)
		reply.Status = computer_iface.ComputerPowerStatus_COMPUTER_POWER_OFF.String()
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_ON:
		err = utils.ComputerOn(po.Address)
		reply.Status = computer_iface.ComputerPowerStatus_COMPUTER_POWER_ON.String()
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_DETECT:
		status := utils.ComputerCheck(po.LanHostname, 3389)
		switch status {
		case computer_iface.ComputerPowerStatus_COMPUTER_POWER_UNKNOWN:
			err = errors.New("检测主机状态失败")
		default:
			reply.Status = status.String()
		}
	}
	if err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, "操作失败, 请检查配置")
		return
	}
	ctl.NewOkResponse(http.StatusOK, reply)
}
