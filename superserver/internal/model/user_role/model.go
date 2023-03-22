package user_role

import "superserver/internal/model"

type UserRoleRef struct {
	model.Base
	UserId int // 用户ID
	RoleId int // 角色ID
}
