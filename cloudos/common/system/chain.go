package system

import (
	"cloudos/common/interface/ecode_iface"
	"context"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const chainContextKey = "chain_message"

var chainContextPool = sync.Pool{New: func() any {
	return new(ChainMessage)
}}

func NewChainMessage() *ChainMessage {
	chain := chainContextPool.Get().(*ChainMessage)
	chain.TraceId = time.Now().UnixNano()
	return chain
}

func GetChainMessage(ctx *gin.Context) *ChainMessage {
	if chain, exists := ctx.Get(chainContextKey); exists {
		return chain.(*ChainMessage)
	}
	return NewChainMessage()
}

func GetChainMessageWithContext(ctx context.Context) *ChainMessage {
	switch value := ctx.Value(chainContextKey).(type) {
	case *ChainMessage:
		return value
	}
	return NewChainMessage()
}

type ChainMessage struct {
	Success bool   `json:"success"`  // 请求成功
	TraceId int64  `json:"trace_id"` // 链路ID
	Message string `json:"message"`  // 异常消息
	Data    any    `json:"data,omitempty"`

	// not export fields
	Level    ecode_iface.ECode `json:"-"`
	UserId   int               `json:"-"`
	AuthTime int64             `json:"-"`
}

func (m *ChainMessage) WriteNormal(data any) {
	m.Success = true
	m.Data = data
}

func (m *ChainMessage) WriteAbnormal(level ecode_iface.ECode, message string) {
	m.Level = level
	m.Message = message
}

func (m *ChainMessage) IntoContext(ctx *gin.Context) {
	ctx.Set(chainContextKey, m)
}

func (m *ChainMessage) GetLevel() ecode_iface.ECode {
	return m.Level
}

func (m *ChainMessage) Recycle() {
	m.TraceId = 0
	m.UserId = 0
	m.Data = nil
	m.Message = ""
	m.Level = ecode_iface.ECode_SUCCESS
}
