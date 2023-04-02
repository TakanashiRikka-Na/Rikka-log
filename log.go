package log

import (
	"go.uber.org/zap"
)

var defaultLogger *zap.SugaredLogger

func init() {
	defaultLogger = NewLogger()
}

func Info(fields ...interface{}) {
	defaultLogger.Info(fields...)
}

func Debug(fields ...interface{}) {
	defaultLogger.Debug(fields...)
}

func Warn(fields ...interface{}) {
	defaultLogger.Warn(fields...)
}

func Error(fields ...interface{}) {
	defaultLogger.Error(fields...)
}

func Fatal(fields ...interface{}) {
	defaultLogger.Fatal(fields...)
}

func Panic(fields ...interface{}) {
	defaultLogger.Panic(fields...)
}

func Infof(format string, fields ...interface{}) {
	defaultLogger.Infof(format, fields...)
}

func Debugf(format string, fields ...interface{}) {
	defaultLogger.Debugf(format, fields...)
}

func Warnf(format string, fields ...interface{}) {
	defaultLogger.Warnf(format, fields...)
}

func Errorf(format string, fields ...interface{}) {
	defaultLogger.Errorf(format, fields...)
}

func Fatalf(format string, fields ...interface{}) {
	defaultLogger.Fatalf(format, fields...)
}

func Panicf(format string, fields ...interface{}) {
	defaultLogger.Panicf(format, fields...)
}
