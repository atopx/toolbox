package auth_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListAccess(ctx context.Context, params *auth_service.ListAccessParams) (*auth_service.ListAccessReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.ListAccess(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.ListAccess failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_ListAccess
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.ListAccess reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateAccess(ctx context.Context, params *auth_service.OperateAccessParams) (*auth_service.OperateAccessReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.OperateAccess(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.OperateAccess failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_OperateAccess
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.OperateAccess reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateAccess(ctx context.Context, params *auth_service.BatchOperateAccessParams) (*auth_service.BatchOperateAccessReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.BatchOperateAccess(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.BatchOperateAccess failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_BatchOperateAccess
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.BatchOperateAccess reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
