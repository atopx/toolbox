package auth_client

import (
	"context"
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"
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
