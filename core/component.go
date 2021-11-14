package core

import (
	"gowebdemo/core/config"
	"gowebdemo/core/logger"
	"gowebdemo/core/mysql"
	"gowebdemo/core/redis"
	"gowebdemo/core/rediscluster"
)

func InitComponent() {
	// 解析配置
	config.SetUp()

	// 初始化日志组件
	logger.SetUp()
	logger.ServerLogger.Info("Logger SetUp Success")

	// 初始化mysql
	mysql.SetUp()
	logger.ServerLogger.Info("Mysql SetUp Success")

	// 初始化redis
	redis.SetUp()
	logger.ServerLogger.Info("Redis SetUp Success")

	// 初始化redis cluter
	rediscluster.SetUp()
	logger.ServerLogger.Info("RedisCluster SetUp Success")

	// 初始化kafka
	// ....

}
