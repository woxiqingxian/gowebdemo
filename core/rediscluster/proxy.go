package rediscluster

import (
	"github.com/go-redis/redis/v8"
)

func InitRedisCluster(name string) *redis.ClusterClient {
	if client, ok := redisClusterClientMap.Load(name); ok {
		if v, ok := client.(*redis.ClusterClient); ok {
			return v
		}
	}
	return nil
}
