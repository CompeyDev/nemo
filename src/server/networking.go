package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CompeyDev/nemo/common/database"
	"github.com/CompeyDev/nemo/common/logger"
	"github.com/gorilla/websocket"
)

// TODO: Use sockets instead of http
// REDO: Fix superfluous calls.

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
		logger.Error("Teamserver failed to launch with error.")
		logger.Error(err.Error())
		os.Exit(1)
	}

}

func heartbeatHandler(writer http.ResponseWriter, request *http.Request) {
	start := time.Now()

	errorResponse := HeartbeatResponse{
		Status: 400,
		Data: map[string]any{
			"uptime":           nil,
			"connectedClients": nil,
		},
	}

	writer.Header().Set("Content-Type", "application/json")

	jsonErrorResponse, errorMarshalErr := json.Marshal(errorResponse)

	if errorMarshalErr != nil {
		log.Fatal("POST /heartbeat -> Failed to encode response, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	if request.Body == nil {
		logger.Error("POST /heartbeat -> No response body provided, Status Code 400")
		writer.WriteHeader(400)
		writer.Write(jsonErrorResponse)
		return
	}

	errorResponse2 := HeartbeatResponse{
		Status: 502,
		Data: map[string]any{
			"uptime":           nil,
			"connectedClients": nil,
		},
	}

	jsonErrorResponse2, errorMarshalErr2 := json.Marshal(errorResponse2)

	if errorMarshalErr2 != nil {
		logger.Error("POST /heartbeat -> Failed to encode response, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse2)
		return
	}

	type RequestBody struct {
		Id string
	}

	var requestBody RequestBody

	jsonParseErr := json.NewDecoder(request.Body).Decode(&requestBody)

	if jsonParseErr != nil {
		logger.Error("POST /heartbeat -> JSON Request parsing error, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse2)
		return
	}

	instanceExists, instanceCheckingErr := database.CheckInstanceExistence(requestBody.Id)

	if instanceCheckingErr != nil {
		logger.Error("Failed to fetch existing payload instances from database, Status Code 502")
		logger.Error(fmt.Sprint(instanceCheckingErr))
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse2)
		return
	}

	if instanceExists {
		database.UpdateCheckInTime(requestBody.Id)
	}

	defer request.Body.Close()

	if jsonParseErr != nil {
		logger.Error("POST /heartbeat -> Failed to encode response, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse2)
		return
	}

	Response := HeartbeatResponse{
		Status: 502,
		Data: map[string]any{
			"uptime":           nil,
			"connectedClients": nil,
		},
	}

	jsonResponse, jsonParseErr := json.Marshal(Response)

	if jsonParseErr != nil {
		logger.Error("POST /heartbeat -> Failed to encode response, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse2)
		return
	}
	writer.WriteHeader(200)
	writer.Write(jsonResponse)
	logger.Info(fmt.Sprintf("POST /heartbeat -> 200 (%s)", time.Since(start).String()))
}
