package logging

import (
	"context"
)

type ctxKey string

const (
	loggerKey = ctxKey("logger")
	depthKey = ctxKey("depth")
)

func NewContext(ctx context.Context, l *Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, loggerKey, l)
}

func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return defaultLogger
	}
	l, ok := ctx.Value(loggerKey).(*Logger)
	if ok {
		return l
	}
	return defaultLogger
}

func withDepth(ctx context.Context, depth int) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, depthKey, depth)
}

func getDepth(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	depth, ok := ctx.Value(depthKey).(int)
	if ok {
		return depth
	}
	return 0
}

func deepen(ctx context.Context) context.Context {
	return withDepth(ctx, getDepth(ctx) + 1)
}

func Deepen(ctx context.Context) context.Context {
	return deepen(ctx)
}
