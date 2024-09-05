package utils

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(msg string, args ...interface{}) {
	infoLogger.Printf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	errorLogger.Printf(msg, args...)
}
