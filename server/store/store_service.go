package store

import (
	"context"
	"fmt"
	"os"
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
func InitaliseStore() *StorageService {
	redisURL := os.Getenv("REDIS_URL") // "redis:6379"
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
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

func SaveUrlMapping(shortUrl string, longUrl string) {
	err := storeService.redisClient.Set(ctx, shortUrl, longUrl, cacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, longUrl))
	}
}

func RetriveInitialUrl(shortUrl string) string {
	res, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		fmt.Printf("No URL found for shortUrl: %s\n", shortUrl)
		return ""
	} else if err != nil {
		panic(fmt.Sprintf("Failed retrieving key url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return res
}
