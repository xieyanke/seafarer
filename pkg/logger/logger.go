package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/xieyanke/seafarer/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(logSetting *setting.LoggerSetting, appSetting *setting.AppSetting) *zap.Logger {
	core := newZapCore(logSetting, appSetting)
	caller := zap.AddCaller()

	opts := zap.Fields(zap.String("application", appSetting.Name))
	logger := zap.New(core, caller, opts)

	return logger
}

func newZapCore(logSetting *setting.LoggerSetting, appSetting *setting.AppSetting) zapcore.Core {
	writer := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logSetting.Dir, logSetting.Filename),
		Compress:   logSetting.Compress,
		MaxSize:    logSetting.MaxSize,
		MaxBackups: logSetting.MaxBackups,
		MaxAge:     logSetting.MaxAge,
	}

	encoder := newZapEncoder(logSetting.Format, appSetting.RunMode)
	level := getLogLevel(logSetting.Level)
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(&writer)), level)

	return core
}

func newZapEncoder(format, runMode string) zapcore.Encoder {
	var cfg zapcore.EncoderConfig
	if runMode == "debug" {
		cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg = zap.NewProductionEncoderConfig()
	}

	cfg.MessageKey = "msg"
	cfg.LevelKey = "level"
	cfg.TimeKey = "time"
	cfg.NameKey = "name"
	cfg.CallerKey = "lines"
	cfg.FunctionKey = "func"
	cfg.StacktraceKey = "stracktrace"
	cfg.LineEnding = zapcore.DefaultLineEnding
	cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeDuration = zapcore.MillisDurationEncoder
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.ConsoleSeparator = " "

	if format == "json" {
		return zapcore.NewJSONEncoder(cfg)
	}

	return zapcore.NewConsoleEncoder(cfg)
}

func getLogLevel(logLevel string) zapcore.Level {
	var level zapcore.Level
	lowercaseLevel := strings.ToLower(logLevel)
	switch lowercaseLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return level
}
