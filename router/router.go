package router

import (
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"web_app/controller"
	"web_app/middleware"
	"web_app/pkg/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"

	_ "web_app/docs" // 千万不要忘了导入把你上一步生成的docs
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
	// 全局中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// v1 group
	v1 := r.Group("/api/v1")
	v1.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, settings.Config.ServerConfig.Name)
	})
	// 不需要 Jwt 认证的路由
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/register", controller.RegisterHandler)

	// 需要 Jwt 认证的路由
	v1.Use(middleware.JwtAuthMiddleware)
	v1.GET("/refreshToken", controller.RefreshTokenHandler)
	v1.GET("/community", controller.CommunityListHandler)
	return r
}
