package redis

import "github.com/go-redis/redis/v8"

// GetRedisConn 获取 redis 链接
func GetRedisConn(name string) *redis.Client {
	if client, ok := redisClientMap.Load(name); ok {
		if v, ok := client.(*redis.Client); ok {
			return v
		}
	}
	return nil
}
