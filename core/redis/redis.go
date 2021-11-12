package redis

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
	redisClientMap sync.Map
)

func SetUp() {
	for _, redisConfig := range config.ServerConfig.RedisConfigList {
		redisConn, err := openConn(redisConfig)
		if err != nil {
			logger.ServerLogger.Panic(fmt.Sprintf("Init Redis Error err:%s", err))
		}
		redisClientMap.LoadOrStore(redisConfig.Name, redisConn)
	}
	return
}

func openConn(redisConfig config.RedisConf) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:               redisConfig.Addr,
		Password:           redisConfig.Password,
		Username:           redisConfig.Username,
		DB:                 redisConfig.Db,
		MaxRetries:         redisConfig.MaxRetries,
		DialTimeout:        time.Duration(redisConfig.DialTimeout) * time.Second,
		ReadTimeout:        time.Duration(redisConfig.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(redisConfig.WriteTimeout) * time.Second,
		PoolSize:           redisConfig.PoolSize,
		MinIdleConns:       redisConfig.MinIdleConns,
		PoolTimeout:        time.Duration(redisConfig.PoolTimeout) * time.Second,
		IdleTimeout:        time.Duration(redisConfig.IdleTimeout) * time.Second,
		MaxConnAge:         time.Duration(redisConfig.MaxConnAge) * time.Second,
		IdleCheckFrequency: time.Duration(redisConfig.IdleCheckFrequency) * time.Second,
	})

	if _, err := redisClient.Ping(context.TODO()).Result(); err != nil {
		return nil, err
	}

	return redisClient, nil
}
