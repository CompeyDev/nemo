package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/CompeyDev/nemo/common/database"
	"github.com/CompeyDev/nemo/common/logger"
)

func QueueHandler(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	start := time.Now()

	internalErrorResponse := HeartbeatResponse{
		Status: 502,
		Data: map[string]any{
			"uptime":           nil,
			"connectedClients": nil,
		},
	}

	writer.Header().Set("Content-Type", "application/json")

	jsonErrorResponse, errorMarshalErr := json.Marshal(internalErrorResponse)

	if errorMarshalErr != nil {
		logger.Error("POST /addQueue -> Failed to encode response, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	type RequestBody struct {
		TaskType  string
		TaskName  string
		PayloadID string
	}

	var requestBody RequestBody

	jsonParseErr := json.NewDecoder(request.Body).Decode(&requestBody)

	if jsonParseErr != nil {
		logger.Error("POST /addQueue -> JSON Request body parsing error, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	dbErr := database.AddToPayloadQueue(requestBody.TaskType, requestBody.PayloadID)

	if dbErr != nil {
		logger.Error("POST /addQueue -> Database write error, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	connectedInstances, fetchErr := database.GetConnectedInstances()

	if fetchErr != nil {
		logger.Error("POST /heartbeat -> Failed fetch connected instances, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	tasksQueue, queueFetchErr := database.GetQueue()

	if queueFetchErr != nil {
		logger.Error("POST /heartbeat -> Failed fetch payload tasks queue, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	successResponse := HeartbeatResponse{
		Status: 200,
		Data: map[string]any{
			"uptime":           nil,
			"connectedClients": connectedInstances,
			"tasksQueue":       tasksQueue,
		},
	}

	jsonSuccessResponse, successMarshalErr := json.Marshal(successResponse)

	if successMarshalErr != nil {
		logger.Error("POST /addQueue -> Failed to encode response, Status Code 502")
		writer.WriteHeader(502)
		writer.Write(jsonErrorResponse)
		return
	}

	writer.WriteHeader(200)
	writer.Write(jsonSuccessResponse)
	logger.Info(fmt.Sprintf("POST /addQueue -> 200 (%s)", time.Since(start).String()))
}
