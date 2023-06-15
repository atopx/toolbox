package system

import (
	"cloudos/common/consts"
	"cloudos/common/pb"
	"cloudos/common/utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    pb.ECode `json:"code"`
	Message string   `json:"message"`
	TraceId string   `json:"traceId"`
	Data    any      `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	resp := &Response{TraceId: utils.NewTraceId()}
	ctx.Set(consts.CK_Response, resp)
	return resp
}

func GetResponse(ctx *gin.Context) *Response {
	if value, ok := ctx.Get(consts.CK_Response); ok {
		return value.(*Response)
	}
	return nil
}
