// Package logger provides a bridge to the logging service.
package logger

import "log"

// LoggerTwo is a implementation of the logger interface.
type LoggerOne struct {
}

func (l *LoggerOne) Log(message string) {
	log.Default().Println("LoggerOne logging message: ", message)
}
