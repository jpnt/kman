package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

/// Usage:
///	log := logger.NewLogger()
///
///	log.Info("This is an info message.")
///	log.Warn("This is an warning message.")
///	log.Error("This is an error message.")

type Logger struct {
	log *logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(logrus.InfoLevel)

	return &Logger{log: log}
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}
