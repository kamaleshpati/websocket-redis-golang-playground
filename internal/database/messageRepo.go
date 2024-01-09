package database

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"github.com/kamaleshpati/wsredisPlayground/internal/utils"
)

var channel = utils.GetEnvironmentVariable("CHANNEL")

func PushMessage(redisClient *redis.Client, msg any) {
	if redisClient == nil {
		panic("redisclient not set")
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	if err := redisClient.Publish(channel, payload).Err(); err != nil {
		panic(err)
	}
	log.Println("succesfully pushed")
}
