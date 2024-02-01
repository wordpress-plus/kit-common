package logx

import (
	"context"
	"go.uber.org/zap/zapcore"
	"time"
)

// A Logger represents a logger.
type Logger interface {
	// Debug logs a message at info level.
	Debug(...any)
	// Debugf logs a message at info level.
	Debugf(string, ...any)
	// Debugv logs a message at info level.
	Debugv(string)
	// Debugw logs a message at info level.
	Debugw(string, ...zapcore.Field)

	// Info logs a message at info level.
	Info(...any)
	// Infof logs a message at info level.
	Infof(string, ...any)
	// Infov logs a message at info level.
	Infov(string)
	// Infow logs a message at info level.
	Infow(string, ...zapcore.Field)

	// Warn logs a message at error level.
	Warn(...any)
	// Warnf logs a message at error level.
	Warnf(string, ...any)
	// Warnv logs a message at error level.
	Warnv(string)
	// Warnw logs a message at error level.
	Warnw(string, ...zapcore.Field)

	// Error logs a message at error level.
	Error(...any)
	// Errorf logs a message at error level.
	Errorf(string, ...any)
	// Errorv logs a message at error level.
	Errorv(string)
	// Errorw logs a message at error level.
	Errorw(string, ...zapcore.Field)

	// WithCallerSkip returns a new logger with the given caller skip.
	WithCallerSkip(skip int) Logger
	// WithContext returns a new logger with the given context.
	WithContext(ctx context.Context) Logger
	// WithDuration returns a new logger with the given duration.
	WithDuration(d time.Duration) Logger
	// WithFields returns a new logger with the given fields.
	WithFields(fields ...zapcore.Field) Logger
}
