package main

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xieyanke/seafarer/docs"
	"go.uber.org/zap"
)

var logger *zap.Logger

// ping godoc
// @Summary ping...
// @Schemes
// @Description do ping
// @Tags Demo
// @Accept json
// @produce json
// @Success 200 {string} pong
// @Router	/ping [get]
func ping(ctx *gin.Context) {
	logger.Info("ping success...")
	ctx.JSON(http.StatusOK, "pong")
}

func init() {
	logger, _ = zap.NewProduction()
}

// @title         Seafarer API
// @version       v0.1.0
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic  BasicAuth
func main() {
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	v1 := r.Group("/api/v1")
	v1.GET("/ping", ping)

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + "8805")
}
