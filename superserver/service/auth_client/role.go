package auth_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListRole(ctx context.Context, params *auth_service.ListRoleParams) (*auth_service.ListRoleReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.ListRole(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.ListRole failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_ListRole
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.ListRole reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateRole(ctx context.Context, params *auth_service.OperateRoleParams) (*auth_service.OperateRoleReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.OperateRole(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.OperateRole failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_OperateRole
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.OperateRole reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateRole(ctx context.Context, params *auth_service.BatchOperateRoleParams) (*auth_service.BatchOperateRoleReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.BatchOperateRole(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.BatchOperateRole failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_BatchOperateRole
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.BatchOperateRole reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
