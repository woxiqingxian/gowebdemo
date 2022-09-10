package rediscluster

import (
	"context"
	"fmt"
	"gowebdemo/core/config"
	"gowebdemo/core/logger"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClusterClientMap sync.Map
)

// SetUp redisCluster 组件初始化
func SetUp() {
	for _, redisClusterConfig := range config.ServerConfig.RedisClusteConfigList {
		redisClusterConn, err := openConn(redisClusterConfig)
		if err != nil {
			logger.ServerLog().Panic(fmt.Sprintf("Init RedisCluster Error err:%s", err))
		}
		redisClusterClientMap.LoadOrStore(redisClusterConfig.Name, redisClusterConn)
		logger.ServerLog().Info(fmt.Sprintf("rediscluster %s setup success", redisClusterConfig.Name))
	}
	return
}

func openConn(redisClusterConfig config.RedisClusterConf) (*redis.ClusterClient, error) {
	redisClusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              redisClusterConfig.Addrs,
		Password:           redisClusterConfig.Password,
		Username:           redisClusterConfig.Username,
		MaxRetries:         redisClusterConfig.MaxRetries,
		RouteByLatency:     redisClusterConfig.RouteByLatency,
		RouteRandomly:      redisClusterConfig.RouteRandomly,
		DialTimeout:        time.Duration(redisClusterConfig.DialTimeout) * time.Second,
		ReadTimeout:        time.Duration(redisClusterConfig.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(redisClusterConfig.WriteTimeout) * time.Second,
		PoolSize:           redisClusterConfig.PoolSize,
		MinIdleConns:       redisClusterConfig.MinIdleConns,
		PoolTimeout:        time.Duration(redisClusterConfig.PoolTimeout) * time.Second,
		IdleTimeout:        time.Duration(redisClusterConfig.IdleTimeout) * time.Second,
		MaxConnAge:         time.Duration(redisClusterConfig.MaxConnAge) * time.Second,
		IdleCheckFrequency: time.Duration(redisClusterConfig.IdleCheckFrequency) * time.Second,
	})

	if _, err := redisClusterClient.Ping(context.TODO()).Result(); err != nil {
		return nil, err
	}

	return redisClusterClient, nil
}
