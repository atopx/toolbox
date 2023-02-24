package create

import (
	"errors"
	"net/http"
	"superserver/internal/model/computer"
	"superserver/proto/computer_iface"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Name == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "名称不能为空")
		return
	}
	if params.Address == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "MAC地址不能为空")
		return
	}

	dao := computer.NewDao(ctl.GetDatabase())
	_, err := dao.FilterBy("name", params.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
			return
		}
		ctl.NewErrorResponse(http.StatusBadRequest, "名称已存在")
		return
	}
	_, err = dao.FilterBy("address", params.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
			return
		}
		ctl.NewErrorResponse(http.StatusBadRequest, "MAC地址已存在")
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

	if err := dao.Create(&cpt); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
		return
	}

	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
