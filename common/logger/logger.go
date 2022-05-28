package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"toolbox/common/system"
)

var logger *zap.Logger

func Setup(level string) (err error) {
	var loggerLevel = new(zapcore.Level)
	if err = loggerLevel.UnmarshalText([]byte(level)); err != nil {
		return err
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}), zapcore.AddSync(os.Stdout), loggerLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	return err
}

func Debug(ctx context.Context, message string, fields ...zapcore.Field) {
	output(ctx, zap.DebugLevel, message, fields...)
}

func Info(ctx context.Context, message string, fields ...zapcore.Field) {
	output(ctx, zap.InfoLevel, message, fields...)
}

func Warn(ctx context.Context, message string, fields ...zapcore.Field) {
	output(ctx, zap.WarnLevel, message, fields...)
}

func Error(ctx context.Context, message string, fields ...zapcore.Field) {
	output(ctx, zap.ErrorLevel, message, fields...)
}

func Fatal(ctx context.Context, message string, fields ...zapcore.Field) {
	output(ctx, zap.FatalLevel, message, fields...)
}

func Panic(ctx context.Context, message string, fields ...zapcore.Field) {
	output(ctx, zap.PanicLevel, message, fields...)
}

func output(ctx context.Context, level zapcore.Level, message string, fields ...zapcore.Field) {
	if entity := logger.Check(level, message); entity != nil {
		if ctx != nil {
			switch value := ctx.Value(system.CONTEXT_KEY).(type) {
			case system.ContextValue:
				fields = append(fields, zap.Object(system.CONTEXT_KEY, value))
			}
		}
		entity.Write(fields...)
	}
}
