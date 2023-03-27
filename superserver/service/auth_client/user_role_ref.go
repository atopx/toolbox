package auth_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListUserRoleRef(ctx context.Context, params *auth_service.ListUserRoleRefParams) (*auth_service.ListUserRoleRefReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.ListUserRoleRef(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.ListUserRoleRef failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_ListUserRoleRef
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.ListUserRoleRef reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateUserRoleRef(ctx context.Context, params *auth_service.OperateUserRoleRefParams) (*auth_service.OperateUserRoleRefReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.OperateUserRoleRef(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.OperateUserRoleRef failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_OperateUserRoleRef
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.OperateUserRoleRef reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateUserRoleRef(ctx context.Context, params *auth_service.BatchOperateUserRoleRefParams) (*auth_service.BatchOperateUserRoleRefReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.BatchOperateUserRoleRef(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.BatchOperateUserRoleRef failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_BatchOperateUserRoleRef
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.BatchOperateUserRoleRef reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
