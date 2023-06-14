package bootstrap

import (
	"github.com/xieyanke/seafarer/config"
	"github.com/xieyanke/seafarer/internal/logger"
)

func InitializeLogger() {
	lg := logger.NewLogger(&config.Config.Logger)
	lg.SetupLogger()
}
