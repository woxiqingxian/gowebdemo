package rediscluster

import (
	"github.com/go-redis/redis/v8"
)

// GetRedisClusterConn 获取 redis 集群链接
func GetRedisClusterConn(name string) *redis.ClusterClient {
	if client, ok := redisClusterClientMap.Load(name); ok {
		if v, ok := client.(*redis.ClusterClient); ok {
			return v
		}
	}
	return nil
}
