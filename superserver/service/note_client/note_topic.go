package note_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/note_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListNoteTopic(ctx context.Context, params *note_service.ListNoteTopicParams) (*note_service.ListNoteTopicReply, ecode.ECode) {
	reply, err := service.GetClient().Note.ListNoteTopic(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.ListNoteTopic failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_ListNoteTopic
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.ListNoteTopic reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateNoteTopic(ctx context.Context, params *note_service.OperateNoteTopicParams) (*note_service.OperateNoteTopicReply, ecode.ECode) {
	reply, err := service.GetClient().Note.OperateNoteTopic(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.OperateNoteTopic failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_OperateNoteTopic
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.OperateNoteTopic reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateNoteTopic(ctx context.Context, params *note_service.BatchOperateNoteTopicParams) (*note_service.BatchOperateNoteTopicReply, ecode.ECode) {
	reply, err := service.GetClient().Note.BatchOperateNoteTopic(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.BatchOperateNoteTopic failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_BatchOperateNoteTopic
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.BatchOperateNoteTopic reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
