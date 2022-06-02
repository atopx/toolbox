package system

import (
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

const CONTEXT_KEY = "server"

var service, version string

func InitContextValue(tags string) {
	index := strings.Index(tags, "-")
	if index > 0 {
		service = tags[:index]
		version = tags[index+1:]
	}
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(ContextValue)
	},
}

func NewContextValue(client string) *ContextValue {
	sys := pool.Get().(*ContextValue)
	sys.TraceId = time.Now().Unix()
	sys.Server = service
	sys.Version = version
	sys.Client = client
	return sys
}

func LoadContextValue(ctx *gin.Context) *ContextValue {
	return ctx.MustGet(CONTEXT_KEY).(*ContextValue)
}

func ContextDone(sys *ContextValue) {
	sys.Client = ""
	sys.UserId = 0
	sys.TraceId = 0
	pool.Put(sys)
}

type ContextValue struct {
	Server  string
	Version string
	UserId  int64
	TraceId int64
	Client  string
}

func (val *ContextValue) SetUserId(userId int64) {
	val.UserId = userId
}

func (val ContextValue) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("service", val.Server)
	encoder.AddString("version", val.Version)
	encoder.AddString("client_ip", val.Client)
	encoder.AddInt64("trace_id", val.TraceId)
	encoder.AddInt64("user_id", val.UserId)
	return nil
}
