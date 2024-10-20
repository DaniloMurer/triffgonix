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

type Logger struct {
	Logger *log.Logger
}

func NewLogger() Logger {
	newLogger := log.New(os.Stdout, "[TRIFFGONIX] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	return Logger{Logger: newLogger}
}

func (logger Logger) Info(message string, args ...any) {
	formattedMessage := fmt.Sprintf(message, args...)
	logger.Logger.Printf("%s: %s\n", INFO, formattedMessage)
}

func (logger Logger) Warn(message string, args ...any) {
	formattedMessage := fmt.Sprintf(message, args...)
	logger.Logger.Printf("%s: %s\n", WARN, formattedMessage)
}

func (logger Logger) Error(message string, args ...any) {
	formattedMessage := fmt.Sprintf(message, args...)
	logger.Logger.Printf("%s: %s\n", ERROR, formattedMessage)
}

func (logger Logger) Trace(message string, args ...any) {
	formattedMessage := fmt.Sprintf(message, args...)

	logger.Logger.Printf("%s: %s\n", WARN, formattedMessage)
}

func (logger Logger) Debug(message string, args ...any) {
	formattedMessage := fmt.Sprintf(message, args...)
	logger.Logger.Printf("%s: %s\n", DEBUG, formattedMessage)
}
