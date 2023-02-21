package system

import (
	"sync"
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

var chainContextPool = sync.Pool{New: func() any {
	return new(ChainContext)
}}

func NewChainContext() *ChainContext {
	chain := chainContextPool.Get().(*ChainContext)
	chain.TraceId = time.Now().UnixNano()
	return chain
}

type ChainContext struct {
	TraceId int64      `json:"trace_id"` // 链路ID
	Message string     `json:"message"`  // 异常消息
	Data    any        `json:"data,omitempty"`
	level   ChainLevel `json:"-"` // not export
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

func (m *ChainContext) Recycle() {
	m.TraceId = 0
	m.Data = nil
	m.Message = ""
	m.level = CHAIN_NORMAL
}
