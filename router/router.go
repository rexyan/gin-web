package router

import (
	"net/http"
	"web_app/controller"
	"web_app/pkg/logger"
	"web_app/settings"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func setRunMode() {
	switch settings.Config.ServerConfig.Mode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.DebugMode:
		gin.SetMode(gin.DebugMode)
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}

func Setup() *gin.Engine {
	setRunMode()
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, viper.Get("server.name"))
	})

	r.POST("/api/v1/register", controller.RegisterHandler)
	return r
}
