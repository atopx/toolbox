package login

import (
	"fmt"
	"net/http"
	"superserver/internal/model/user"
	"time"
)

func (ctl *Controller) Deal() {
	userPo := &user.User{Username: ctl.param.Username}
	userPo.SetPassword(ctl.param.Password)
	userDao := user.NewDao(ctl.GetDatabase())
	if err := userDao.Login(userPo); err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, "invalid username or password")
		return
	}
	if userPo.Status == user.UserStatusDisabled.Enum {
		ctl.NewErrorResponse(http.StatusBadRequest, "invalid username or password")
		return
	}

	// TODO set cookie. session

	// update last login ip and time
	userPo.LastLoginTime = time.Now().Unix()
	userPo.SetLastLoginIp(ctl.Context.ClientIP())
	if err := userDao.Update(userPo); err != nil {
		ctl.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("update user last login info failed: %s", err.Error()))
		return
	}
	reply := &Reply{
		Id:              userPo.Id,
		Username:        userPo.Username,
		Role:            user.UserRoleEnumMap[userPo.Role],
		AvatarUrl:       userPo.AvatarUrl,
		Status:          user.UserStatusEnumMap[userPo.Status],
		TotalSizeLimit:  userPo.TotalSizeLimit,
		SingleSizeLimit: userPo.SingleSizeLimit,
		SizeUsed:        userPo.SizeUsed,
		LastLoginIp:     userPo.GetLastLoginIp(),
		LastLoginTime:   userPo.LastLoginTime,
		CreateTime:      userPo.CreateTime,
		UpdateTime:      userPo.UpdateTime,
	}
	ctl.NewOkResponse(http.StatusOK, reply)
}
