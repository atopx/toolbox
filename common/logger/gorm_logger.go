package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

var DefaultGormLoggerConfig = gormLogger.Config{
	IgnoreRecordNotFoundError: false,
	SlowThreshold:             200 * time.Millisecond,
	LogLevel:                  InfoLevel,
}

const (
	SilentLevel gormLogger.LogLevel = iota + 1
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

var LogLevel_value = map[string]gormLogger.LogLevel{
	"debug":  DebugLevel,
	"info":   InfoLevel,
	"warn":   WarnLevel,
	"error":  ErrorLevel,
	"silent": SilentLevel,
}

type GormLogger struct {
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	Colorful                            bool
	gormLogger.Config
}

func NewGormLogger() gormLogger.Interface {
	var (
		infoStr      = "%s [info] "
		warnStr      = "%s [warn] "
		errStr       = "%s [error] "
		traceStr     = "%s [%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
	)

	gormLogger := &GormLogger{
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
	gormLogger.Config = DefaultGormLoggerConfig
	return gormLogger
}

func (l *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel <= InfoLevel {
		logStr := fmt.Sprintf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		Info(ctx, logStr)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= WarnLevel {
		logStr := fmt.Sprintf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		Warn(ctx, logStr)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= ErrorLevel {
		logStr := fmt.Sprintf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
		Error(ctx, logStr)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= SilentLevel {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.LogLevel >= ErrorLevel && (!errors.Is(err, gormLogger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		Error(ctx, fmt.Sprintf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= WarnLevel:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		Warn(ctx, fmt.Sprintf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
	case l.LogLevel == InfoLevel:
		Info(ctx, fmt.Sprintf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
	}
}
