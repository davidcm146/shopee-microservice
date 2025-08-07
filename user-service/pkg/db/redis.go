package db

import (
	"context"
	"fmt"
	"log"

	"github.com/davidcm146/shopee-microservice/user-service/config"
	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Client *redis.Client
	Ctx    context.Context
}

func InitRedis() *RedisConfig {
	addr := config.LoadEnv("REDIS_ADDR", "localhost:6379")

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connect error:", err)
	}
	fmt.Println("Connected to Redis")

	return &RedisConfig{
		Client: client,
		Ctx:    ctx,
	}
}
