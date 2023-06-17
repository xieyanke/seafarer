package global

import (
	"github.com/xieyanke/seafarer/pkg/setting"
	"go.uber.org/zap"
)

var (
	ServerSetting *setting.ServerSetting
	AppSetting    *setting.AppSetting
	LoggerSetting *setting.LoggerSetting
	SqliteSetting *setting.SqliteSetting
)

var (
	Logger *zap.Logger
)
