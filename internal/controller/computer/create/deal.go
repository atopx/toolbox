package create

import (
	"errors"
	"net/http"
	"superserver/internal/model/computer"
	"superserver/internal/model/user"
	"superserver/proto/computer_iface"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	if ctl.param.Name == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "名称不能为空")
		return
	}
	if ctl.param.Address == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "MAC地址不能为空")
		return
	}

	dao := computer.NewDao(ctl.GetDatabase())

	if !errors.Is(dao.FilterBy("name", ctl.param.Name), gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "名称已存在")
		return
	}

	if !errors.Is(dao.FilterBy("address", ctl.param.Name), gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "MAC地址已存在")
		return
	}

	cpt := computer.Computer{
		Name:        ctl.param.Name,
		Username:    ctl.param.Username,
		Password:    ctl.param.Password,
		LanHostname: ctl.param.LanHostname,
		WanHostname: ctl.param.WanHostname,
		Address:     ctl.param.Address,
		PowerStatus: computer_iface.ComputerPowerStatus_COMPUTER_POWER_UNKNOWN,
		// TODO 待修改
		Creator: user.SystemUser.Id,
		Updator: user.SystemUser.Id,
	}
	if err := dao.Create(&cpt); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "创建主机失败, 请联系管理员")
		return
	}

	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
