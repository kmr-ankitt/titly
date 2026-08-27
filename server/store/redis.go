package store

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisURL = func() string {
	if url := os.Getenv("REDIS_URL"); url != "" {
		return url
	}
	return "localhost:6379"
}()

func InitaliseRedisStore() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	return rdb
}

func GetLongUrlFromRedisStore(ctx context.Context, rdb *redis.Client, shortUrl string) (string, error) {
	longUrl, err := rdb.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		panic(err)
	} 

	return longUrl, nil
}

func PutUrlMappingInRedisStore(ctx context.Context, rdb *redis.Client, shortUrl string, longUrl string) error {
	err := rdb.Set(ctx, shortUrl, longUrl, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}