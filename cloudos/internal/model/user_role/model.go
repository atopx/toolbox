package user_role

import "cloudos/internal/model"

type UserRoleRef struct {
	model.Base
	UserId int // 用户ID
	RoleId int // 角色ID
}
