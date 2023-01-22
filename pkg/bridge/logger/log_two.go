// Package logger provides a bridge to the logging service.
package logger

import "log"

// LoggerTwo is a implementation of the logger interface.
type LoggerTwo struct {
}

func (l *LoggerTwo) Log(message string) {
	log.Default().Println("LoggerTwo logging message: ", message)
}
