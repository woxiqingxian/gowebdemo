package core

import (
	"fmt"
	"gowebdemo/core/config"
	"gowebdemo/core/logger"
	"gowebdemo/core/middleware"
	"gowebdemo/router"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func RunHttpServer(shutdownCallbackFuncList ...func()) {
	// 服务关闭回调
	shutdownBySignal := func() {
		// 给业务的回调
		for _, fc := range shutdownCallbackFuncList {
			fc()
		}
		// 日志flush
		logger.LoggerSync()
	}

	mode := gin.ReleaseMode
	if config.ServerConfig.AppConfig.Debug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	gin.DisableConsoleColor()

	ginEngine := gin.New()

	// 中间件的顺序不能变
	// 引入traceId
	ginEngine.Use(middleware.Trace())
	// 使用日志中间件zap
	if config.ServerConfig.AppConfig.AccessLog {
		ginEngine.Use(logger.Ginzap())
	}
	// recover掉项目可能出现的panic
	if !config.ServerConfig.AppConfig.Debug {
		ginEngine.Use(logger.RecoveryWithZap())
	}

	// 加载路由
	router.Register(ginEngine)

	// 启动服务器(使用endless)
	logger.ServerLogger.Info("server started success")
	addr := fmt.Sprintf("%s:%d", config.ServerConfig.AppConfig.HttpAddr, config.ServerConfig.AppConfig.HttpPort)
	httpServer := endless.NewServer(addr, ginEngine)
	httpServer.SignalHooks[endless.PRE_SIGNAL][syscall.SIGTERM] = append(
		httpServer.SignalHooks[endless.PRE_SIGNAL][syscall.SIGTERM],
		shutdownBySignal, // 服务关闭回调
	)
	httpServer.SignalHooks[endless.PRE_SIGNAL][syscall.SIGINT] = append(
		httpServer.SignalHooks[endless.PRE_SIGNAL][syscall.SIGINT],
		shutdownBySignal, // 服务关闭回调
	)

	err := httpServer.ListenAndServe()
	if err != nil {
		logger.ServerLogger.Error(fmt.Sprintf("server start failed, error: %s", err.Error()))
	}
}
