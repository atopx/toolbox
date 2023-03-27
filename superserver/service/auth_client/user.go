package auth_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListUser(ctx context.Context, params *auth_service.ListUserParams) (*auth_service.ListUserReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.ListUser(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.ListUser failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_ListUser
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.ListUser reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateUser(ctx context.Context, params *auth_service.OperateUserParams) (*auth_service.OperateUserReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.OperateUser(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.OperateUser failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_OperateUser
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.OperateUser reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateUser(ctx context.Context, params *auth_service.BatchOperateUserParams) (*auth_service.BatchOperateUserReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.BatchOperateUser(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.BatchOperateUser failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_BatchOperateUser
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.BatchOperateUser reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
