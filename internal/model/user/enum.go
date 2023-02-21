package user

import "superserver/internal/model"

var (
	UserStatusNormal   = model.Enum{Enum: "NORMAL", Value: "正常"}
	UserStatusDisabled = model.Enum{Enum: "DISABLE", Value: "禁用"}

	UserStatusEnumMap = map[string]model.Enum{
		UserStatusNormal.Enum:   UserStatusNormal,
		UserStatusDisabled.Enum: UserStatusDisabled,
	}

	UserStatusValueMap = map[string]model.Enum{
		UserStatusNormal.Value:   UserStatusNormal,
		UserStatusDisabled.Value: UserStatusDisabled,
	}
)

var (
	UserRoleGuest = model.Enum{Enum: "GUEST", Value: "游客"}
	UserRoleUser  = model.Enum{Enum: "USER", Value: "普通用户"}
	UserRoleAdmin = model.Enum{Enum: "ADMIN", Value: "管理员"}

	UserRoleEnumMap = map[string]model.Enum{
		UserRoleGuest.Enum: UserRoleGuest,
		UserRoleUser.Enum:  UserRoleUser,
		UserRoleAdmin.Enum: UserRoleAdmin,
	}

	UserRoleValueMap = map[string]model.Enum{
		UserRoleGuest.Value: UserRoleGuest,
		UserRoleUser.Value:  UserRoleUser,
		UserRoleAdmin.Value: UserRoleAdmin,
	}
)
