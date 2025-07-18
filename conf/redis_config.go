package conf

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewRedisClient() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		redisClient: client,
		ctx:         context.Background(),
	}, nil
}

func (rc *RedisClient) Set(shortUrl string, url string) error {
	return nil
}

func (rc *RedisClient) Get(shortUrl string) (string, error) {
	return "", nil
}
