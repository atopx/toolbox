package create

import (
	"errors"
	"go.uber.org/zap"
	"superserver/common/interface/computer_iface"
	"superserver/common/interface/ecode_iface"
	"superserver/common/logger"
	"superserver/internal/model/computer"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Name == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "名称不能为空")
		return
	}
	if params.Address == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "MAC地址不能为空")
		return
	}

	dao := computer.NewDao(ctl.GetDatabase())
	_, err := dao.FilterBy("name", params.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			logger.Error(ctl.Context, "create computer dao.FilterByName failed", zap.Error(err))
			ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
			return
		}
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "名称已存在")
		return
	}
	_, err = dao.FilterBy("address", params.Address)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			logger.Error(ctl.Context, "create computer dao.FilterByAddress failed", zap.Error(err))
			ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
			return
		}
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "MAC地址已存在")
		return
	}

	cpt := computer.Computer{
		Name:        params.Name,
		Username:    params.Username,
		Password:    params.Password,
		LanHostname: params.LanHostname,
		WanHostname: params.WanHostname,
		PowerStatus: computer_iface.ComputerPowerStatus_COMPUTER_POWER_UNKNOWN,
		Creator:     ctl.GetOperator(),
		Updater:     ctl.GetOperator(),
	}
	cpt.SetAddress(params.Address)
	if err = dao.Create(&cpt); err != nil {
		logger.Error(ctl.Context, "create computer dao.Create failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		return
	}

	ctl.NewOkResponse(&Reply{})
}
