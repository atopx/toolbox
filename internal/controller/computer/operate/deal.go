package operate

import (
	"errors"
	"fmt"
	"net/http"
	"superserver/internal/model/computer"
	"superserver/proto/computer_iface"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Operate == nil || params.Id <= 0 {
		ctl.NewErrorResponse(http.StatusBadRequest, "无效的操作")
		return
	}

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
		fmt.Println("// TODO 关机")
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_ON:
		fmt.Println("// TODO 开机")
	case computer_iface.ComputerOperate_COMPUTER_OPERATE_DETECT:
		fmt.Println("// TODO 探测")
	}

	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
