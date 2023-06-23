package router

import (
	"time"

	ginZap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xieyanke/seafarer/docs"
	v1 "github.com/xieyanke/seafarer/pkg/api/v1"
	"go.uber.org/zap"
)

var (
	apiVersion = "/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(ginZap.Ginzap(zap.L(), time.RFC3339, false))
	r.Use(ginZap.RecoveryWithZap(zap.L(), true))

	apiv1 := r.Group(apiVersion)

	{
		pingPtr := v1.NewPing()
		apiv1.GET("/ping", pingPtr.Get)
	}

	docs.SwaggerInfo.BasePath = apiVersion
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
