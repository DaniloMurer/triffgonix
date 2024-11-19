package logging

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to capture log output for the custom Logger
func captureLoggerOutput(logger Logger, logFunc func()) string {
	var buf bytes.Buffer
	logger.Logger.SetOutput(&buf)
	defer logger.Logger.SetOutput(os.Stdout) // Restore to default after test

	logFunc()

	return buf.String()
}

// Test NewLogger
func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger.Logger == nil {
		t.Errorf("expected logger to be initialized, got nil")
	}
}

// Test Info logging
func TestLogger_Info(t *testing.T) {
	logger := NewLogger()
	message := captureLoggerOutput(logger, func() {
		logger.Info("test info log")
	})
	if message == "" {
		t.Error("expected log message to be written, got empty string")
	} else if !containsLogLevel(message, "INFO") {
		t.Errorf("expected log message to contain 'INFO', got: %s", message)
	}
}

// Test Warn logging
func TestLogger_Warn(t *testing.T) {
	logger := NewLogger()
	message := captureLoggerOutput(logger, func() {
		logger.Warn("test warn log")
	})
	if message == "" {
		t.Error("expected log message to be written, got empty string")
	} else if !containsLogLevel(message, "WARNING") {
		t.Errorf("expected log message to contain 'WARNING', got: %s", message)
	}
}

// Test Error logging
func TestLogger_Error(t *testing.T) {
	logger := NewLogger()
	message := captureLoggerOutput(logger, func() {
		logger.Error("test error log")
	})
	if message == "" {
		t.Error("expected log message to be written, got empty string")
	} else if !containsLogLevel(message, "ERROR") {
		t.Errorf("expected log message to contain 'ERROR', got: %s", message)
	}
}

// Test Trace logging (when VERBOSE is set)
func TestLogger_Trace(t *testing.T) {
	logger := NewLogger()
	err := os.Setenv("VERBOSE", "true")
	if err != nil {
		return
	}
	message := captureLoggerOutput(logger, func() {
		logger.Trace("test trace log")
	})
	if message == "" {
		t.Error("expected log message to be written, got empty string")
	} else if !containsLogLevel(message, "TRACE") {
		t.Errorf("expected log message to contain 'TRACE', got: %s", message)
	}
	err = os.Unsetenv("VERBOSE")
	if err != nil {
		return
	}
}

// Test Debug logging
func TestLogger_Debug(t *testing.T) {
	logger := NewLogger()
	message := captureLoggerOutput(logger, func() {
		logger.Debug("test debug log")
	})
	if message == "" {
		t.Error("expected log message to be written, got empty string")
	} else if !containsLogLevel(message, "DEBUG") {
		t.Errorf("expected log message to contain 'DEBUG', got: %s", message)
	}
}

// Helper function to check if log message contains the correct log level
func containsLogLevel(message, level string) bool {
	return bytes.Contains([]byte(message), []byte(level))
}
