package system

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

const CONTEXT_KEY = "server"

var (
	service, version string
)

func init() {
	service = viper.GetString("server.name")
	version = os.Getenv("VERSION")
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(ContextValue)
	},
}

func NewContextValue() *ContextValue {
	sys := pool.Get().(*ContextValue)
	sys.TraceId = time.Now().Unix()
	sys.Service = service
	sys.Version = version
	return sys
}

func LoadContextValue(ctx *gin.Context) *ContextValue {
	return ctx.MustGet(CONTEXT_KEY).(*ContextValue)
}

func ContextDone(sys *ContextValue) {
	sys.UserId = 0
	sys.TraceId = 0
	pool.Put(sys)
}

type ContextValue struct {
	Service string
	Version string
	UserId  int64
	TraceId int64
}

func (val *ContextValue) SetUserId(userId int64) {
	val.UserId = userId
}

func (val ContextValue) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("service", val.Service)
	encoder.AddString("version", val.Version)
	encoder.AddInt64("trace_id", val.TraceId)
	encoder.AddInt64("user_id", val.TraceId)
	return nil
}
