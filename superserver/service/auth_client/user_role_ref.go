package auth_client

import (
	"context"
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"
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
