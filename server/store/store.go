package store

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type StoreService struct {
	SqliteClient *sql.DB
	RedisClient *redis.Client
}

func InitialiseStore () *StoreService{
	return &StoreService {
		SqliteClient: InitaliseSqliteStore(),
		RedisClient: InitaliseRedisStore(),
	}
}