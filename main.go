package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/pkg/logger"
	"web_app/pkg/mysql"
	"web_app/pkg/redis"
	"web_app/pkg/snowflake"
	"web_app/pkg/validator"
	"web_app/router"
	"web_app/settings"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

// @title Gin Web
// @version 1.0
// @description Gin Web Project
// @termsOfService http://yanrs.me
// @contact.name runsha.yan
// @contact.url http://yanrs.me
// @contact.email rex_yan@126.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8081
// @BasePath /api/v1/
func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("settings init failed, err:%v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Config.LoggerConfig); err != nil {
		fmt.Printf("logger init failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()

	// 3. 初始化 MySQL
	if err := mysql.Init(settings.Config.MySQLConfig); err != nil {
		fmt.Printf("mysql init failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化 Redis
	if err := redis.Init(settings.Config.RedisConfig); err != nil {
		fmt.Printf("redis init failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// snowflake ID 生成器
	if err := snowflake.Init(); err != nil {
		fmt.Printf("snowflake init failed, err:%v\n", err)
		return
	}

	// 错误翻译
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	// 5. 注册路由
	r := router.Setup()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("server.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
