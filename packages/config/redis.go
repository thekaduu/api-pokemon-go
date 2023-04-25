package config

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func Get(key string) (string, error) {
	ctx := context.Background()
	client := getRedisClient()

	value, err := client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}

func Set(key string, value interface{}) {
	ctx := context.Background()
	client := getRedisClient()

	err := client.Set(ctx, key, value, 5*time.Minute).Err()

	if err != nil {
		panic(err)
	}
}
