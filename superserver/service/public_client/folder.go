package public_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/public/ecode"
	"superserver/domain/public_service"
	"superserver/service"

	"go.uber.org/zap"
)

func ListFolder(ctx context.Context, params *public_service.ListFolderParams) (*public_service.ListFolderReply, ecode.ECode) {
	reply, err := service.GetClient().Public.ListFolder(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.ListFolder failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_ListFolder
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.ListFolder reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateFolder(ctx context.Context, params *public_service.OperateFolderParams) (*public_service.OperateFolderReply, ecode.ECode) {
	reply, err := service.GetClient().Public.OperateFolder(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.OperateFolder failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_OperateFolder
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.OperateFolder reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateFolder(ctx context.Context, params *public_service.BatchOperateFolderParams) (*public_service.BatchOperateFolderReply, ecode.ECode) {
	reply, err := service.GetClient().Public.BatchOperateFolder(ctx, params)
	if err != nil {
		logger.Error(ctx, "public_service.BatchOperateFolder failed", zap.Error(err))
		return nil, ecode.ECode_PUBLIC_SERVICE_ERROR_BatchOperateFolder
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "public_service.BatchOperateFolder reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
