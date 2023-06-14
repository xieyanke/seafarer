package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/xieyanke/seafarer/config"
	"github.com/xieyanke/seafarer/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	cfg *config.LoggerConfig
}

func NewLogger(cfg *config.LoggerConfig) *Logger {
	return &Logger{cfg}
}

func (l *Logger) SetupLogger() error {

	err := util.TryToMkdir(l.cfg.Dir, os.ModePerm)

	writeSyncer := l.getLogWriter()
	encoder := l.getLogEncoder()
	level := l.getLogLevel()
	core := zapcore.NewCore(encoder, writeSyncer, level)
	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)

	return err
}

func (l *Logger) getLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	if strings.ToLower(l.cfg.Format) == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *Logger) getLogLevel() zapcore.Level {
	var level zapcore.Level
	lowercaseLevel := strings.ToLower(l.cfg.Level)
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
		// options = append(options, zap.AddStacktrace(level))
	default:
		level = zap.InfoLevel
		// options = append(options, zap.AddStacktrace(level))
	}

	return level
}

func (l *Logger) getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", l.cfg.Dir, l.cfg.Filename),
		MaxSize:    l.cfg.MaxSize,
		MaxAge:     l.cfg.MaxDays,
		MaxBackups: l.cfg.MaxBackups,
		LocalTime:  false,
		Compress:   l.cfg.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
