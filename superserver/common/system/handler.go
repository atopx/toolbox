package system

import (
	"context"
	"github.com/gin-gonic/gin"
	"superserver/common/consts"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
)

type Response struct {
	Header *common.ReplyHeader `json:"header"`
	Data   any                 `json:"data"`
}

func NewServiceHeader(header *common.Header) *common.Header {
	return &common.Header{
		TraceId:  header.TraceId,
		Source:   consts.ServiceName,
		Operator: header.Operator,
	}
}

func SetRequestHeader(ctx *gin.Context, header *common.Header) {
	ctx.Set(consts.RequestHeader, header)
}

func SetResponseHeader(ctx *gin.Context, header *common.ReplyHeader) {
	ctx.Set(consts.ResponseHeader, header)
}

func GetRequestHeader(ctx context.Context) *common.Header {
	if header, ok := ctx.Value(consts.RequestHeader).(*common.Header); ok {
		return header
	}
	return nil
}

func GetResponseHeader(ctx *gin.Context) *common.ReplyHeader {
	if value, exists := ctx.Get(consts.ResponseHeader); exists {
		if header, ok := value.(*common.ReplyHeader); ok {
			return header
		}
	}
	return nil
}

var User *auth_service.User
