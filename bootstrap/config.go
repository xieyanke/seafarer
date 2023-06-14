package bootstrap

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xieyanke/seafarer/config"
)

var configFile = "config.yaml"

func InitializeConfig() *viper.Viper {
	if configEnv := os.Getenv("SEAFARER_CONFIG"); configEnv != "" {
		configFile = configEnv
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %v", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&config.Config); err != nil {
			panic(fmt.Errorf("reload config failed: %v", err))
		}
	})

	if err := v.Unmarshal(&config.Config); err != nil {
		panic(fmt.Errorf("read config failed: %v", err))
	}

	return v
}
