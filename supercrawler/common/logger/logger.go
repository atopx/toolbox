package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}), zapcore.AddSync(os.Stdout), loggerLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	return err
}

func GetLogger() *zap.Logger {
	return logger
}

func AddCallerSkip(n int) *zap.Logger {
	return logger.WithOptions(zap.AddCallerSkip(n))
}

func System(message string, v ...any) {
	logger.Info(fmt.Sprintf(message, v...))
}

func Debug(message string, fields ...zapcore.Field) {
	output(zap.DebugLevel, message, fields...)
}

func Info(message string, fields ...zapcore.Field) {
	fmt.Println(message)
	output(zap.InfoLevel, message, fields...)
}

func Warn(message string, fields ...zapcore.Field) {
	output(zap.WarnLevel, message, fields...)
}

func Error(message string, fields ...zapcore.Field) {
	output(zap.ErrorLevel, message, fields...)
}

func Fatal(message string, fields ...zapcore.Field) {
	output(zap.FatalLevel, message, fields...)
}

func Panic(message string, fields ...zapcore.Field) {
	output(zap.PanicLevel, message, fields...)
}

func output(level zapcore.Level, message string, fields ...zapcore.Field) {
	if entity := logger.Check(level, message); entity != nil {
		entity.Write(fields...)
	}
}
