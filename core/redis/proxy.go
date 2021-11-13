package redis

import "github.com/go-redis/redis/v8"

func InitRedis(name string) *redis.Client {
	if client, ok := redisClientMap.Load(name); ok {
		if v, ok := client.(*redis.Client); ok {
			return v
		}
	}
	return nil
}
