package bootstrap

import (
	"github.com/xieyanke/seafarer/global"
	"github.com/xieyanke/seafarer/pkg/logger"
	"go.uber.org/zap"
)

func SetupLogger() {
	l := logger.NewLogger(global.LoggerSetting, global.ServerSetting, global.AppSetting)
	zap.ReplaceGlobals(l)
}
