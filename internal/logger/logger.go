package logger

import (
	"context"

	"go.uber.org/zap"
)

type ContextKey struct{}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	logger, ok := getLoggerFromContext(ctx)
	if !ok {
		return getNewLogger()
	}
	return logger
}

func GetNewContextWithLogger() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKey{}, getNewLogger())

	FromContext(ctx).Info("Logger successfully created for the context")

	return ctx
}

func getNewLogger() *zap.SugaredLogger {
	// Adopted from Zap's Quick Start https://github.com/uber-go/zap
	l, _ := zap.NewProduction()
	defer l.Sync()
	logger := l.Sugar()

	return logger
}

func getLoggerFromContext(ctx context.Context) (*zap.SugaredLogger, bool) {
	l, ok := ctx.Value(ContextKey{}).(*zap.SugaredLogger)
	return l, ok
}
