package logger

import (
	"fmt"
	"log"
	"os"
)

type Level string

const (
	INFO  Level = "INFO"
	WARN  Level = "WARNING"
	ERROR Level = "ERROR"
	TRACE Level = "TRACE"
	DEBUG Level = "DEBUG"
)

// TODO: set desired logging level via env variable. e.g no trace logs when production
type Logger struct {
	Logger *log.Logger
}

func NewLogger() Logger {
	newLogger := log.New(os.Stdout, "[TRIFFGONIX] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	return Logger{Logger: newLogger}
}

func (logger Logger) Info(message string, args ...any) {
	var formattedMessage string
	if args != nil {
		formattedMessage = fmt.Sprintf(message, args...)
	} else {
		formattedMessage = message
	}
	logger.Logger.Printf("%s: %s\n", INFO, formattedMessage)
}

func (logger Logger) Warn(message string, args ...any) {
	var formattedMessage string
	if args != nil {
		formattedMessage = fmt.Sprintf(message, args...)
	} else {
		formattedMessage = message
	}
	logger.Logger.Printf("%s: %s\n", WARN, formattedMessage)
}

func (logger Logger) Error(message string, args ...any) {
	var formattedMessage string
	if args != nil {
		formattedMessage = fmt.Sprintf(message, args...)
	} else {
		formattedMessage = message
	}
	logger.Logger.Printf("%s: %s\n", ERROR, formattedMessage)
}

func (logger Logger) Trace(message string, args ...any) {
	var formattedMessage string
	if args != nil {
		formattedMessage = fmt.Sprintf(message, args...)
	} else {
		formattedMessage = message
	}
	logger.Logger.Printf("%s: %s\n", TRACE, formattedMessage)
}

func (logger Logger) Debug(message string, args ...any) {
	var formattedMessage string
	if args != nil {
		formattedMessage = fmt.Sprintf(message, args...)
	} else {
		formattedMessage = message
	}
	logger.Logger.Printf("%s: %s\n", DEBUG, formattedMessage)
}