package conf

import (
	"context"
	"fmt"
	"log"

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

func (rc *RedisClient) Set(shortUrl string, url string) {
	fmt.Printf("-- initiating SET: %s, %s --\n", shortUrl, url)
	err := rc.redisClient.Set(rc.ctx, shortUrl, url, 0).Err()
	if err != nil {
		log.Panicf("failed to SET in Redis: %s\n", err)
		return
	}
	fmt.Println("-- complete SET --")
}

func (rc *RedisClient) Get(shortUrl string) string {
	fmt.Printf("-- initiating GET: %s --\n", shortUrl)
	url, err := rc.redisClient.Get(rc.ctx, shortUrl).Result()
	if err != nil {
		log.Panicf("failed to GET in Redis: %s\n", err)
	}
	fmt.Println("-- complete GET --")
	return url
}
