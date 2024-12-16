package ctxmanage

import (
	"book-store/middlewares"
	"context"
	"log/slog"
)

func GetTraceID(ctx context.Context) string {
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		slog.Error("traceId not found in the context")
		return "UNKNOWN"
	}
	return traceId
}
