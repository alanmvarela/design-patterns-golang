// Package logger provides a bridge to the logging service.
package logger

import (
	"testing"
)

// Test LoggerOne tests the logger
func TestLoggerOne(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("LoggerOne panicked: %s", r)
		}
	}()
	logger := LoggerOne{}
	// Send a log to the logger
	logger.Log("example.com")
}

// Test LoggerTwo tests the logger
func TestLoggerTwo(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("LoggerTwo panicked: %s", r)
		}
	}()
	logger := LoggerTwo{}
	// Send a log to the logger
	logger.Log("example.com")
}
