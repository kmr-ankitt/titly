package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)


func InitaliseRedisStore() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
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