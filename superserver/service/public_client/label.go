package public_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service"

	"go.uber.org/zap"
)

func ListLabel(ctx context.Context, params *public_service.ListLabelParams) (*public_service.ListLabelReply, ecode.ECode) {
	reply, err := service.GetClient().Public.ListLabel(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.ListLabel failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_ListLabel
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.ListLabel reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateLabel(ctx context.Context, params *public_service.OperateLabelParams) (*public_service.OperateLabelReply, ecode.ECode) {
	reply, err := service.GetClient().Public.OperateLabel(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.OperateLabel failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_OperateLabel
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.OperateLabel reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateLabel(ctx context.Context, params *public_service.BatchOperateLabelParams) (*public_service.BatchOperateLabelReply, ecode.ECode) {
	reply, err := service.GetClient().Public.BatchOperateLabel(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.BatchOperateLabel failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_BatchOperateLabel
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.BatchOperateLabel reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
