package auth_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListPermission(ctx context.Context, params *auth_service.ListPermissionParams) (*auth_service.ListPermissionReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.ListPermission(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.ListPermission failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_ListPermission
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.ListPermission reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperatePermission(ctx context.Context, params *auth_service.OperatePermissionParams) (*auth_service.OperatePermissionReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.OperatePermission(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.OperatePermission failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_OperatePermission
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.OperatePermission reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperatePermission(ctx context.Context, params *auth_service.BatchOperatePermissionParams) (*auth_service.BatchOperatePermissionReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.BatchOperatePermission(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.BatchOperatePermission failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_BatchOperatePermission
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.BatchOperatePermission reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
