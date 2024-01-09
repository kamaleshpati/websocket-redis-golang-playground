package database

import (
	"sync"

	"github.com/go-redis/redis"

	"github.com/kamaleshpati/wsredisPlayground/internal/utils"
)

var clientInstance *redis.Client
var lock = &sync.Mutex{}

func GetRedisClient() *redis.Client {
	return getSingleInstance(clientInstance)
}

func getSingleInstance(singleInstance *redis.Client) *redis.Client {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = createInstance()
		}
	}

	return singleInstance
}

func createInstance() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.GetEnvironmentVariable("DBHOST"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
