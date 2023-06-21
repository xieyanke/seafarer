package main

import (
	"github.com/xieyanke/seafarer/bootstrap"
	"github.com/xieyanke/seafarer/global"
	"github.com/xieyanke/seafarer/internal/router"
)

func init() {
	bootstrap.SetupSetting()
	bootstrap.SetupLogger()
}

// @title         Seafarer API
// @version       v0.1.0
// @schemes       http
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic  BasicAuth
func main() {
	router := router.NewRouter()

	router.Run(":" + global.ServerSetting.HttpPort)
}
