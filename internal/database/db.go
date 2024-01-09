package database

import (
	"github.com/go-redis/redis"

	"github.com/kamaleshpati/wsredisPlayground/internal/utils"
)

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.GetEnvironmentVariable("DBHOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
