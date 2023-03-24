package info

import (
	"errors"
	"superserver/common/interface/ecode_iface"
	"superserver/common/logger"
	"superserver/internal/model"
	"superserver/internal/model/user"
	"superserver/internal/model/user_role"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	userPo := user.User{Base: model.Base{Id: ctl.GetOperator()}}
	dao := user.NewDao(ctl.GetDatabase())
	err := dao.Load(&userPo)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "用户不存在")
		return
	}
	if err != nil {
		logger.Error(ctl.Context, "user info dao.Load failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "[12001]系统错误, 请联系管理员")
		return
	}
	reply := Reply{
		Id:         userPo.Id,
		Name:       userPo.Name,
		Username:   userPo.Username,
		LoginTime:  ctl.GetAuthTime(),
		CreateTime: userPo.CreateTime,
		UpdateTime: userPo.UpdateTime,
		Roles:      []Role{},
	}

	// get user roles
	roles, err := user_role.NewDao(ctl.GetDatabase()).GetRoles(userPo.Id)
	if err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "用户不存在")
		return
	}
	for _, role := range roles {
		reply.Roles = append(reply.Roles, Role{
			Id:   role.Id,
			Name: role.Name,
		})
	}

	ctl.NewOkResponse(&reply)
}
