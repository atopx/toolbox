package create

import (
	"cloudos/common/interface/ecode_iface"
	"cloudos/common/interface/user_iface"
	"cloudos/common/logger"
	"cloudos/internal/model/role"
	"cloudos/internal/model/user"
	"cloudos/internal/model/user_role"
	"errors"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Name == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_INVALID_PARAMS, "名称不能为空")
		return
	}
	if params.Username == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_INVALID_PARAMS, "用户名不能为空")
		return
	}
	if params.Password == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_INVALID_PARAMS, "密码不能为空")
		return
	}
	dao := user.NewDao(ctl.GetDatabase())

	_, err := dao.GetUserByUsername(params.Username, true)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_REQUEST_CONFLICT, "用户名已存在")
		return
	}

	if err != nil {
		logger.Error(ctl.Context, "create user dao.GetUserByUsername failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_DB_ERROR, "系统错误, 请联系管理员")
		return
	}

	po := user.User{
		Name:     params.Name,
		Username: params.Username,
		Status:   user_iface.UserStatus_USER_NORMAL,
	}
	po.SetPassword(params.Password)

	if err = dao.Create(&po); err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_DB_ERROR, "系统错误, 请联系管理员")
		return
	}

	// 绑定默认角色
	if err = user_role.NewDao(ctl.GetDatabase()).Create(&user_role.UserRoleRef{UserId: po.Id, RoleId: role.DefaultRole.Id}); err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_DB_ERROR, "系统错误, 请联系管理员")
		return
	}

	ctl.NewOkResponse(&Reply{})
}
