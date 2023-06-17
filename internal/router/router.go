package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xieyanke/seafarer/docs"
	v1 "github.com/xieyanke/seafarer/internal/api/v1"
)

var (
	apiVersion = "/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group(apiVersion)

	{
		pingPtr := v1.NewPing()
		apiv1.GET("/ping", pingPtr.Get)
	}

	docs.SwaggerInfo.BasePath = apiVersion
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
