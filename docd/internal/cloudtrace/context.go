package cloudtrace

import (
	"context"
)

type traceKey struct{}
type spanKey struct{}

func contextWithTraceInfo(ctx context.Context, traceHeader string) context.Context {
	traceID, spanID := parseHeader(traceHeader)
	if traceID != "" {
		ctx = context.WithValue(ctx, traceKey{}, traceID)
	}
	if spanID != "" {
		ctx = context.WithValue(ctx, spanKey{}, spanID)
	}
	return ctx
}

func traceInfoFromContext(ctx context.Context) (traceID, spanID string) {
	traceID, _ = ctx.Value(traceKey{}).(string)
	spanID, _ = ctx.Value(spanKey{}).(string)
	return
}
