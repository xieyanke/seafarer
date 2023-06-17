package bootstrap

import (
	"path/filepath"
	"time"

	"github.com/xieyanke/seafarer/global"
	"github.com/xieyanke/seafarer/pkg/setting"
)

var configFile = "config/config.yaml"

func SetupSetting() error {
	setting, err := setting.NewSetting(filepath.Dir(configFile), filepath.Base(configFile), "yaml")
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Logger", &global.LoggerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Sqlite", &global.SqliteSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.ReadHeaderTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
