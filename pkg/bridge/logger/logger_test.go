// Package logger provides a bridge to the logging service.
package logger

import (
	"testing"
)

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
