package system

import (
	"time"
)

type ChainLevel int8

const (
	CHAIN_NORMAL ChainLevel = iota
	CHAIN_BAD
	CHAIN_ERROR
)

const (
	ChainContextKey = "chain"
)

type ChainContext struct {
	TraceId int64  `json:"trace_id"` // 链路ID
	Message string `json:"message"`  // 异常消息
	Data    any    `json:"data,omitempty"`

	level ChainLevel
}

func NewChainContext() *ChainContext {
	return &ChainContext{TraceId: time.Now().UnixNano()}
}

func (m *ChainContext) WriteNormal(data any) {
	m.Data = data
}

func (m *ChainContext) WriteAbnomal(level ChainLevel, message string) {
	m.level = level
	m.Message = message

}

func (m *ChainContext) GetLevel() ChainLevel {
	return m.level
}
