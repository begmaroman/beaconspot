package logex

import (
	"fmt"

	log "github.com/go-log/log"
	"go.uber.org/zap"
)

// microLogger implements log.Logger interface to provide custom logger for go-micro packages.
type microLogger struct {
	logger *zap.Logger
}

// NewMicro is the constructor of microLogger.
func NewMicro(logger *zap.Logger) log.Logger {
	return &microLogger{
		logger: logger,
	}
}

// Log implements log.Logger interface.
func (l *microLogger) Log(v ...interface{}) {
	fields := make([]zap.Field, len(v))
	for i, e := range v {
		fields[i] = zap.Any("", e)
	}
	l.logger.Info("", fields...)
}

// Logf implements log.Logger interface.
func (l *microLogger) Logf(format string, v ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, v...))
}
