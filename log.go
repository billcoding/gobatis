package gobatis

import (
	"fmt"
	l "log"
)

// log struct
type log struct {
	ologger *l.Logger
	elogger *l.Logger
}

// Info level
func (l *log) Info(message string, args ...interface{}) {
	l.ologger.Println(fmt.Sprintf(message, args...))
}

// Warn level
func (l *log) Warn(message string, args ...interface{}) {
	l.ologger.Println(fmt.Sprintf(message, args...))
}

// Error level
func (l *log) Error(message string, args ...interface{}) {
	l.elogger.Println(fmt.Sprintf(message, args...))
}
