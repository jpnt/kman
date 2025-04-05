package logger

import (
	"fmt"
	"sync"
	"time"
)

type ILogger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

type Logger struct {
	level LogLevel
	mu    sync.Mutex
}

var _ ILogger = (*Logger)(nil)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	WarnLevel
	ErrorLevel
)

func NewLogger(l LogLevel) *Logger {
	return &Logger{level: l}
}

func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= InfoLevel {
		l.log("INFO", "\033[32m", format, args...)
	}
}

func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= WarnLevel {
		l.log("WARN", "\033[33m", format, args...)
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ErrorLevel {
		l.log("ERRO", "\033[31m", format, args...)
	}
}

func (l *Logger) log(level, color, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	timestamp := time.Now().Format("2006-01-02T15:04:05.00000")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s%s\033[0m: %s\n", timestamp, color, level, message)
}
