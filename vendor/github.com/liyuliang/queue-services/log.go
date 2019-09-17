package services

import (
	"log"
	"fmt"
)

var isDebug bool = true

func Debug(format string, msg ...interface{}) {
	if isDebug {
		log.Println(fmt.Sprintf(format, msg...))
	}
}

func Info(format string, msg ...interface{}) {
	Debug(format, msg...)
}

func Error(format string, msg ...interface{}) {
	Debug(format, msg...)
}
