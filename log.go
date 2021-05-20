package gobatis

import (
	"fmt"
	l "log"
)

// log struct
type log struct {
	outLogger *l.Logger
	errLogger *l.Logger
}

// Info level
func (l *log) Info(message string, args ...interface{}) {
	l.outLogger.Println(fmt.Sprintf(message, args...))
}

// Warn level
func (l *log) Warn(message string, args ...interface{}) {
	l.outLogger.Println(fmt.Sprintf(message, args...))
}

// Error level
func (l *log) Error(message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	l.errLogger.Println(msg)
	panic(msg)
}
