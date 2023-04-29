package model

import (
	"superserver/common/enum"
	"superserver/common/utils"
	"superserver/domain/auth_service"
)

type Role struct {
	Id         int32                   `json:"id"`
	Name       string                  `json:"name"`
	Nature     auth_service.RoleNature `json:"nature"`
	NatureDesc string                  `json:"nature_desc"`
	CreateTime string                  `json:"create_time"`
	UpdateTime string                  `json:"update_time"`
}

func NewRole(src *auth_service.Role) Role {
	return Role{
		Id:         src.Id,
		Name:       src.Name,
		Nature:     src.Nature,
		NatureDesc: enum.Desc(enum.RoleNature, int32(src.Nature)),
		CreateTime: utils.TimestampDefaultDecoder(src.CreateTime),
		UpdateTime: utils.TimestampDefaultDecoder(src.UpdateTime),
	}
}

func NewRoleList(list []*auth_service.Role) []Role {
	result := make([]Role, 0, len(list))
	for _, src := range list {
		result = append(result, NewRole(src))
	}
	return result
}
