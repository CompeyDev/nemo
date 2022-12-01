package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CompeyDev/nemo/common/logger"
	"github.com/gorilla/websocket"
)

// TODO: Use sockets instead of http

type HeartbeatResponse struct {
	Status int            `json:"status"`
	Data   map[string]any `json:"data"`
}

var Registry = map[string]func(http.ResponseWriter, *http.Request){
	"heartbeat": heartbeatHandler,
}

func GetHandler(toFetchHandler string) func(http.ResponseWriter, *http.Request) {
	handler, checkValidHandler := Registry[toFetchHandler]
	if checkValidHandler {
		return handler
	} else if !checkValidHandler {
		return nil
	}

	return nil
}

var wss = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func toWebsocket(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	ws, err := wss.Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println("Failed to upgrade websocket connection.")
	}

	return ws
}

func InitializeConnection() {

	for endpoint, handler := range Registry {
		http.HandleFunc(fmt.Sprintf("/%s", endpoint), handler)
	}

	logger.CustomInfo("ready", "Launched teamserver on wss://0.0.0.0:40043")

	if err := http.ListenAndServe("0.0.0.0:40043", nil); err != nil {
		logger.Error("Teamserver failed to launch.")
		os.Exit(1)
	}

}

func heartbeatHandler(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()
	defer request.Body.Close()
	writer.Header().Set("Content-Type", "application/json")

	response := HeartbeatResponse{
		Status: 200,
		Data: map[string]any{
			"uptime":           nil,
			"connectedClients": nil,
		},
	}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatal("Failed to encode response!")
	}

	logger.Info(fmt.Sprintf("GET /heartbeat -> 200 (%s)", time.Since(start).String()))

	writer.Write(jsonResponse)
}
