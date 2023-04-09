package note_client

import (
	"context"
	"superserver/common/logger"
	"superserver/domain/note_service"
	"superserver/domain/public/ecode"
	"superserver/service"

	"go.uber.org/zap"
)

func ListNote(ctx context.Context, params *note_service.ListNoteParams) (*note_service.ListNoteReply, ecode.ECode) {
	reply, err := service.GetClient().Note.ListNote(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.ListNote failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_ListNote
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.ListNote reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func OperateNote(ctx context.Context, params *note_service.OperateNoteParams) (*note_service.OperateNoteReply, ecode.ECode) {
	reply, err := service.GetClient().Note.OperateNote(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.OperateNote failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_OperateNote
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.OperateNote reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}

func BatchOperateNote(ctx context.Context, params *note_service.BatchOperateNoteParams) (*note_service.BatchOperateNoteReply, ecode.ECode) {
	reply, err := service.GetClient().Note.BatchOperateNote(ctx, params)
	if err != nil {
		logger.Error(ctx, "note_service.BatchOperateNote failed", zap.Error(err))
		return nil, ecode.ECode_NOTE_SERVICE_ERROR_BatchOperateNote
	}
	if reply.Header.Code != ecode.ECode_SUCCESS {
		logger.Error(ctx, "note_service.BatchOperateNote reply error", zap.String("error", reply.Header.Message))
	}
	return reply, reply.Header.Code
}
