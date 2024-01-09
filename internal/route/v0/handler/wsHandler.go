package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"

	"github.com/kamaleshpati/wsredisPlayground/internal/database"
	"github.com/kamaleshpati/wsredisPlayground/internal/route/v0/model"
)

const (
	connected    = "CONNECTED"
	recieved     = "RECIVED"
	send         = "SEND"
	wrongMessage = "ERROR"
)

var connection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var wsOnMsgRecieveChannel = make(chan model.WsJsonPayload)

var redisClient *redis.Client

// starting the pre needed
func init() {
	redisClient = database.GetRedisClient()
	go processWsMessageService()
}

func WSEndpointHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := connection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client is connected to server:wsEndpoint")

	var response model.WsJsonResponse
	response.Action = connected
	response.Message = "connected to ws-server successfully."

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go listenToWsChannel(ws)
}

// the goroutine that listens for our channels
func listenToWsChannel(wsConnection *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	var response model.WsJsonResponse
	var payload model.WsJsonPayload
	for {
		err := wsConnection.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			response.Action = wrongMessage
			response.Message = err.Error()
			wsConnection.WriteJSON(response)
		} else {
			wsOnMsgRecieveChannel <- payload
			response.Action = recieved
			response.Message = "message recieved sucessfully"
			wsConnection.WriteJSON(response)
		}
	}
}

// listens to channel and pushes data
func processWsMessageService() {
	for {
		msg := <-wsOnMsgRecieveChannel

		switch msg.Action {
		case send:
			database.PushMessage(redisClient, msg)
		default:
			log.Println("Not an proper message")
		}
	}
}
