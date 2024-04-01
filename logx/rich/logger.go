package rich

import (
	"context"
	"go.uber.org/zap/zapcore"
	"time"
)

// A Logger represents a logger.
type Logger interface {
	// Debuga logs a message at info level.
	Debuga(...any)
	// Debugf logs a message at info level.
	Debugf(string, ...any)
	// Debugv logs a message at info level.
	Debugv(string)
	// Debug logs a message at info level.
	Debug(string, ...zapcore.Field)

	// Infoa logs a message at info level.
	Infoa(...any)
	// Infof logs a message at info level.
	Infof(string, ...any)
	// Infov logs a message at info level.
	Infov(string)
	// Info logs a message at info level.
	Info(string, ...zapcore.Field)

	// Warna logs a message at error level.
	Warna(...any)
	// Warnf logs a message at error level.
	Warnf(string, ...any)
	// Warnv logs a message at error level.
	Warnv(string)
	// Warn logs a message at error level.
	Warn(string, ...zapcore.Field)

	// Errora logs a message at error level.
	Errora(...any)
	// Errorf logs a message at error level.
	Errorf(string, ...any)
	// Errorv logs a message at error level.
	Errorv(string)
	// Errorw logs a message at error level.
	Error(string, ...zapcore.Field)

	// WithCallerSkip returns a new logger with the given caller skip.
	WithCallerSkip(skip int) Logger
	// WithContext returns a new logger with the given context.
	WithContext(ctx context.Context) Logger
	// WithDuration returns a new logger with the given duration.
	WithDuration(d time.Duration) Logger
	// WithFields returns a new logger with the given fields.
	WithFields(fields ...zapcore.Field) Logger
}
