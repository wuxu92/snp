package utils

import (
	"github.com/astaxie/beego/logs"
	// "time"
	"fmt"
)

var logger *logs.BeeLogger
var console *logs.BeeLogger

func GetLogger() *logs.BeeLogger {
	if logger == nil {
		logger = logs.NewLogger(1000)
		logger.SetLevel(logs.LevelInfo)
//		today := time.Now().YearDay()
		logger.SetLogger("file", `{"filename":"logs/snp.log","daily":true}`)
	}
	return logger
}

func GetConsole() *logs.BeeLogger {
	if console == nil {
		fmt.Println("new console logger")
		console = logs.NewLogger(1000)
		console.SetLevel(logs.LevelInfo)
		console.SetLogger("console", `{}`)
	}
	return console
}