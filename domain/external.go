package domain

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type External interface {
	DB() *gorm.DB
	RedisClient() *redis.Client
}
