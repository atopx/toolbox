package auth_client

import (
	"context"
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/domain/auth_service"
	"superserver/domain/public/ecode"
	"superserver/service"
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
