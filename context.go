package logging

import (
	"context"
)

type ctxKey string

func NewContext(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, ctxKey("logger"), l)
}

func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return defaultLogger
	}
	l, ok := ctx.Value(ctxKey("logger")).(*Logger)
	if ok {
		return l
	}
	return defaultLogger
}
