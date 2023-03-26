package auth_client

import (
	"context"
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"
)

func ListAuthToken(ctx context.Context, params *auth_service.ListAuthTokenParams) (*auth_service.ListAuthTokenReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.ListAuthToken(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.ListAuthToken failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_ListAuthToken
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.ListAuthToken reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateAuthToken(ctx context.Context, params *auth_service.OperateAuthTokenParams) (*auth_service.OperateAuthTokenReply, ecode.ECode) {
	reply, err := service.GetClient().Auth.OperateAuthToken(ctx, params)
	if err != nil {
		logger.Error(ctx, "auth_service.OperateAuthToken failed", zap.Error(err))
		return nil, ecode.ECode_AUTH_SERVICE_ERROR_OperateAuthToken
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "auth_service.OperateAuthToken reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
