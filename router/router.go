package router

import (
	"net/http"
	"web_app/controller"
	"web_app/middleware"
	"web_app/pkg/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func setRunMode() {
	// 设置运行模式。默认值: debug, 支持 release, test, debug
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
		context.JSON(http.StatusOK, settings.Config.ServerConfig.Name)
	})

	r.POST("/api/v1/login", controller.LoginHandler)
	r.POST("/api/v1/register", controller.RegisterHandler)
	r.GET("/api/v1/refreshToken", middleware.JwtAuthMiddleware, controller.RefreshTokenHandler)
	return r
}
