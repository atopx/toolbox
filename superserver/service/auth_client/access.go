package auth_client

import (
	"context"
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"
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
