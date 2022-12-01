package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
)

var timeStyle = color.Gray.Print
var infoStyle = color.Green.Print
var errorStyle = color.Red.Print

func getLogTime() string {
	return strings.Split(strings.Split(time.Now().String(), "m=")[0], "+")[0]
}

func Info(log string) {
	timeStyle(getLogTime())
	infoStyle("info")
	fmt.Print(" :: ")
	fmt.Println(log)
}

func Error(log string) {
	timeStyle(getLogTime())
	errorStyle("error")
	fmt.Print(" :: ")
	fmt.Println(log)
}

func CustomInfo(caller string, log string) {
	timeStyle(getLogTime())
	infoStyle(caller)
	fmt.Print(" :: ")
	fmt.Println(log)
}

func CustomError(caller string, log string) {
	timeStyle(getLogTime())
	errorStyle(caller)
	fmt.Print(" :: ")
	fmt.Println(log)
}
