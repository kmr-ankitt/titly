package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func initaliseStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6359",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis Started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}
