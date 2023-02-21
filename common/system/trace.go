package system

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Span struct {
	Id        int64 `json:"id"`
	Localtion string
}

type Trace struct {
	Id    int64  `json:"id"`
	Spans []Span `json:"span"`
}

func NewTrace() *Trace {
	return &Trace{Id: time.Now().UnixNano(), Spans: make([]Span, 0)}
}

const (
	TRACE_KEY = "super_trace"
)

func (t Trace) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddInt64("traceid", t.Id)
	// TODO add spans
	return nil
}
