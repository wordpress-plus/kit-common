package rich

import (
	"context"
	"go.uber.org/zap/zapcore"
	"sync"
	"sync/atomic"
)

var (
	fieldsContextKey contextKey
	globalFields     atomic.Value
	globalFieldsLock sync.Mutex
)

type contextKey struct{}

// AddGlobalFields adds global fields.
func AddGlobalFields(fields ...zapcore.Field) {
	globalFieldsLock.Lock()
	defer globalFieldsLock.Unlock()

	old := globalFields.Load()
	if old == nil {
		globalFields.Store(append([]zapcore.Field(nil), fields...))
	} else {
		globalFields.Store(append(old.([]zapcore.Field), fields...))
	}
}

// ContextWithFields returns a new context with the given fields.
func ContextWithFields(ctx context.Context, fields ...zapcore.Field) context.Context {
	if val := ctx.Value(fieldsContextKey); val != nil {
		if arr, ok := val.([]zapcore.Field); ok {
			allFields := make([]zapcore.Field, 0, len(arr)+len(fields))
			allFields = append(allFields, arr...)
			allFields = append(allFields, fields...)
			return context.WithValue(ctx, fieldsContextKey, allFields)
		}
	}

	return context.WithValue(ctx, fieldsContextKey, fields)
}
