package create

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/user"
	"superserver/proto/user_iface"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Name == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "名称不能为空")
		return
	}
	if params.Username == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "用户名不能为空")
		return
	}
	if params.Password == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "密码不能为空")
		return
	}
	dao := user.NewDao(ctl.GetDatabase())

	_, err := dao.GetUserByUsername(params.Username, true)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			logger.Error(ctl.Context, "create user dao.GetUserByUsername failed", zap.Error(err))
			ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
			return
		}
		ctl.NewErrorResponse(http.StatusBadRequest, "用户名已存在")
		return
	}

	po := user.User{
		Name:     params.Name,
		Username: params.Username,
		//RoleId:   rolePo.Id,
		Status: user_iface.UserStatus_USER_NORMAL,
	}
	po.SetPassword(params.Password)

	if err := dao.Create(&po); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(http.StatusOK, &Reply{})
}
