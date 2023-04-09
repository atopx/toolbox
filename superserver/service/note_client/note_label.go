package note_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/note_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListNoteLabel(ctx context.Context, params *note_service.ListNoteLabelParams) (*note_service.ListNoteLabelReply, ecode.ECode) {
	reply, err := service.GetClient().Note.ListNoteLabel(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.ListNoteLabel failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_ListNoteLabel
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.ListNoteLabel reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateNoteLabel(ctx context.Context, params *note_service.OperateNoteLabelParams) (*note_service.OperateNoteLabelReply, ecode.ECode) {
	reply, err := service.GetClient().Note.OperateNoteLabel(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.OperateNoteLabel failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_OperateNoteLabel
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.OperateNoteLabel reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateNoteLabel(ctx context.Context, params *note_service.BatchOperateNoteLabelParams) (*note_service.BatchOperateNoteLabelReply, ecode.ECode) {
	reply, err := service.GetClient().Note.BatchOperateNoteLabel(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.BatchOperateNoteLabel failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_BatchOperateNoteLabel
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.BatchOperateNoteLabel reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
