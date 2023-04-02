package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() *zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()
	// 时间编码调整
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// log 等级首字母大写
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// console 风格的 Encoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	// 标准输出控制台输出的 writer
	consoleWriter := zapcore.AddSync(os.Stdout)
	// log 文件的输出的 writer
	logfileWriter := zapcore.AddSync(GetIOLogger("./log.log"))

	// default 控制台输出的等级
	defaultConsoleLogLevel := zapcore.DebugLevel
	// default log 文件输出的等级
	defaultFileLogLevel := zapcore.InfoLevel

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleWriter, defaultConsoleLogLevel),
		zapcore.NewCore(consoleEncoder, logfileWriter, defaultFileLogLevel),
	)

	// 添加调用者，以及当是 error 等级的时候打出调用栈
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func GetIOLogger(logFilePath string) io.Writer {
	return &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   false,
	}
}
