package external

import (
	"github.com/go-redis/redis/v8"
	"uv-chat-api-server-golang/internal/config"
)

func mustRedis(config *config.Config) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 서버 주소
	})
	return c
}
