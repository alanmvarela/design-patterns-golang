// Package logger provides a bridge to the logging service.
package logger

type Logger interface {
	// Log logs the given message.
	Log(message string)
}
