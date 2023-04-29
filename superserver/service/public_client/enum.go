package public_client

import (
	"context"
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service"
)

func ListEnum(ctx context.Context, params *public_service.ListEnumParams) (*public_service.ListEnumReply, ecode.ECode) {
	reply, err := service.GetClient().Public.ListEnum(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.ListEnum failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_ListEnum
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.ListEnum reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
