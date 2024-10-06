package logger

import (
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

func (logger *Logger) NewLogger() Logger {
	newLogger := log.New(os.Stdout, "[TRIFFGONIX] - ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	return Logger{Logger: newLogger}
}

func (logger *Logger) Info(message string) {
	logger.Logger.Printf("%s: %s", INFO, message)
}

func (logger *Logger) Warn(message string) {
	logger.Logger.Printf("%s: %s", WARN, message)
}

func (logger *Logger) Error(message string) {
	logger.Logger.Printf("%s: %s", ERROR, message)
}

func (logger *Logger) Trace(message string) {
	logger.Logger.Printf("%s: %s", WARN, message)
}

func (logger *Logger) Debug(message string) {
	logger.Logger.Printf("%s: %s", DEBUG, message)
}
