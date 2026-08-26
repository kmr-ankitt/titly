package store

import (
	"context"
	"fmt"

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

func TestRedisConnection(ctx context.Context, rdb *redis.Client) {
	err := rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val) // >>> foo bar

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
	err := rdb.Set(ctx, shortUrl, longUrl, 604800).Err()
	if err != nil {
		return err
	}

	return nil
}

//TODO: Just for testing
func GetAllMappingsFromRedisStore(ctx context.Context, rdb *redis.Client) (map[string]string, error) {
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	mappings := make(map[string]string)
	for _, key := range keys {
		longUrl, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		mappings[key] = longUrl
	}

	return mappings, nil
}