package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kamaleshpati/wsredisPlayground/internal/database"
	"github.com/kamaleshpati/wsredisPlayground/internal/route"
	"github.com/kamaleshpati/wsredisPlayground/internal/route/v0/model"
	"github.com/kamaleshpati/wsredisPlayground/internal/utils"
)

var host = utils.GetEnvironmentVariable("HOST")
var channel = utils.GetEnvironmentVariable("CHANNEL")

func main() {

	go PullMessages()

	mux := route.Routes()

	// start the web server
	log.Println("Starting app on port:", host)
	err := http.ListenAndServe(host, mux)
	if err != nil {
		freehost, err := utils.GetFreePort()
		if err != nil {
			panic(err)
		}
		log.Println("Starting app on port:", freehost)
		http.ListenAndServe(freehost, mux)
	}
}

func PullMessages() {
	redisClient := database.GetRedisClient()
	subscriber := redisClient.Subscribe(channel)
	obj := model.WsJsonPayload{}
	for {
		msg, err := subscriber.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &obj); err != nil {
			panic(err)
		}

		log.Println("Received message from " + msg.Channel + " channel.")
		log.Printf("%+v\n", obj)
	}
}
