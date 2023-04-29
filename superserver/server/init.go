package server

import (
	"context"
	"superserver/common/consts"
	"superserver/common/enum"
	"superserver/common/system"
	"superserver/common/utils"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service/auth_client"
	"superserver/service/public_client"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func initAccess(ctx context.Context, routes gin.RoutesInfo) {
	var authAccessMap = map[string]auth_service.AccessStatus{
		"/api/user/login":   auth_service.AccessStatus_ACCESS_ANONYMOUS,
		"/api/user/refresh": auth_service.AccessStatus_ACCESS_ANONYMOUS,
	}
	var accessList []*auth_service.Access
	for _, routeInfo := range routes {
		record := &auth_service.Access{
			Path:   routeInfo.Path,
			Method: routeInfo.Method,
			Status: authAccessMap[routeInfo.Path],
		}
		accessList = append(accessList, record)
	}
	_, code := auth_client.BatchOperateAccess(ctx, &auth_service.BatchOperateAccessParams{
		Header:  &common.Header{TraceId: -1, Source: consts.ServiceName},
		Operate: common.Operation_OPERATION_UPSERT,
		Data:    accessList,
	})
	if code != ecode.ECode_SUCCESS {
		panic(system.GetErrorMessage(code))
	}
}

func initSystemUser(ctx context.Context) {
	user := &auth_service.User{
		Username: viper.GetString("admin.user"),
		Password: utils.Hash(viper.GetString("admin.pass")),
		Status:   auth_service.UserStatus_USER_NORMAL,
	}
	listUserReply, code := auth_client.ListUser(ctx, &auth_service.ListUserParams{
		Header: &common.Header{TraceId: -1, Source: consts.ServiceName},
		Pager:  &common.Pager{Disabled: true},
		Filter: &auth_service.UserFilter{Usernames: []string{user.Username}},
	})
	if code != ecode.ECode_SUCCESS {
		panic(system.GetErrorMessage(code))
	}
	if len(listUserReply.Data) == 0 {
		// create
		operateUserReply, code := auth_client.OperateUser(ctx, &auth_service.OperateUserParams{
			Header:  &common.Header{TraceId: -1, Source: consts.ServiceName},
			Operate: common.Operation_OPERATION_CREATE,
			Data:    user,
		})
		if code != ecode.ECode_SUCCESS {
			panic(system.GetErrorMessage(code))
		}
		system.User = &auth_service.User{
			Id:       operateUserReply.Data.Id,
			Username: user.Username,
		}
	} else {
		system.User = &auth_service.User{
			Id:       listUserReply.Data[0].Id,
			Username: user.Username,
		}
	}
}

func initInternalRoles(ctx context.Context) {
	header := &common.Header{TraceId: -1, Source: consts.ServiceName, Operator: system.User.Id}
	roles := []*auth_service.Role{
		{Name: "系统管理员", Nature: auth_service.RoleNature_ROLE_SYSTEM},
		{Name: "初始角色", Nature: auth_service.RoleNature_ROLE_DEFAULT},
	}
	for _, role := range roles {
		listRoleReply, code := auth_client.ListRole(ctx, &auth_service.ListRoleParams{
			Header: header, Pager: &common.Pager{Disabled: true},
			Filter: &auth_service.RoleFilter{Natures: []auth_service.RoleNature{role.Nature}},
		})
		if code != ecode.ECode_SUCCESS {
			panic(system.GetErrorMessage(code))
		}
		if len(listRoleReply.Data) == 0 {
			// create
			operateRoleReply, code := auth_client.OperateRole(ctx, &auth_service.OperateRoleParams{
				Header:  &common.Header{TraceId: -1, Source: consts.ServiceName},
				Operate: common.Operation_OPERATION_CREATE,
				Data:    role,
			})
			if code != ecode.ECode_SUCCESS {
				panic(system.GetErrorMessage(code))
			}
			role.Id = operateRoleReply.Data.Id
		} else {
			role.Id = listRoleReply.Data[0].Id
		}
	}

	// bind system user and system role
	listUserRoleRefReply, code := auth_client.ListUserRoleRef(ctx, &auth_service.ListUserRoleRefParams{
		Header: header, Pager: &common.Pager{Disabled: true},
		Filter: &auth_service.UserRoleRefFilter{RoleIds: []int32{roles[0].Id}, UserIds: []int32{system.User.Id}},
	})
	if code != ecode.ECode_SUCCESS {
		panic(system.GetErrorMessage(code))
	}
	if len(listUserRoleRefReply.Data) == 0 {
		// create
		_, code := auth_client.OperateUserRoleRef(ctx, &auth_service.OperateUserRoleRefParams{
			Header: header, Operate: common.Operation_OPERATION_CREATE,
			Data: &auth_service.UserRoleRef{
				UserId: system.User.Id,
				RoleId: roles[0].Id,
			},
		})
		if code != ecode.ECode_SUCCESS {
			panic(system.GetErrorMessage(code))
		}
	}
}

func initEnums(ctx context.Context) {
	header := &common.Header{TraceId: -1, Source: consts.ServiceName, Operator: system.User.Id}
	reply, code := public_client.ListEnum(ctx, &public_service.ListEnumParams{Header: header})
	if code != ecode.ECode_SUCCESS {
		panic(system.GetErrorMessage(code))
	}
	enum.InitEnums(reply.Data)
}
