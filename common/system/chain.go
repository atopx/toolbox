package system

import (
	"context"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ChainLevel int8

const (
	CHAIN_NORMAL ChainLevel = iota
	CHAIN_BAD
	CHAIN_ERROR
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
	TraceId int64  `json:"trace_id"` // 链路ID
	Message string `json:"message"`  // 异常消息
	Data    any    `json:"data,omitempty"`

	// not export fields
	level  ChainLevel `json:"-"`
	UserId int        `json:"-"`
}

func (m *ChainMessage) WriteNormal(data any) {
	m.Data = data
}

func (m *ChainMessage) WriteAbnomal(level ChainLevel, message string) {
	m.level = level
	m.Message = message
}

func (m *ChainMessage) IntoContext(ctx *gin.Context) {
	ctx.Set(chainContextKey, m)
}

func (m *ChainMessage) GetLevel() ChainLevel {
	return m.level
}

func (m *ChainMessage) Recycle() {
	m.TraceId = 0
	m.UserId = 0
	m.Data = nil
	m.Message = ""
	m.level = CHAIN_NORMAL
}
