package log

import (
	"SenseHoney/lib/utils"
	"fmt"
	"log"
	"os"
)

func WriteLog(level string, content interface{}) {
	var fileName string
	switch level {
	case "error":
		fileName = "error.log"
	// info 包括 warn
	case "info":
		fileName = "info.log"
	case "debug":
		fileName = "debug.log"
	default:
		panic("日志等级错误")
	}

	logFile, err := os.OpenFile("./logs/"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //windows需先创建logs目录
	if err != nil {
		fmt.Println("Write logFile error: ", err.Error())
		os.Exit(1)
	} else {
		logger := log.New(logFile, "[SenseHoney] ", log.Lshortfile|log.Ldate|log.Ltime)
		logger.Printf("%v", content)
	}

}

var apiLogUrl = "http://" + utils.GetConfig().API.APIAddr + "/api/log"

// 日志信息上报格式
type Log struct {
	Level       string `json:"level"`
	AccessToken string `json:"accessToken"`
}
