package cloudtrace

// Inspired by https://github.com/remko/cloudrun-slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// Extra log level supported by Cloud Logging
const LevelCritical = slog.Level(12)

// Handler that outputs JSON understood by the structured log agent.
// See https://cloud.google.com/logging/docs/agent/logging/configuration#special-fields
type CloudLoggingHandler struct {
	handler   slog.Handler
	projectID string
}

var _ slog.Handler = (*CloudLoggingHandler)(nil)

func NewCloudLoggingHandler(projectID string, level slog.Level) *CloudLoggingHandler {
	return &CloudLoggingHandler{
		projectID: projectID,
		handler: slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.MessageKey {
					a.Key = "message"
				} else if a.Key == slog.SourceKey {
					a.Key = "logging.googleapis.com/sourceLocation"
				} else if a.Key == slog.LevelKey {
					a.Key = "severity"
					level := a.Value.Any().(slog.Level)
					if level == LevelCritical {
						a.Value = slog.StringValue("CRITICAL")
					}
				}
				return a
			},
		}),
	}
}

func (h *CloudLoggingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *CloudLoggingHandler) Handle(ctx context.Context, rec slog.Record) error {
	traceID, spanID := traceInfoFromContext(ctx)
	if traceID != "" {
		rec = rec.Clone()
		// https://cloud.google.com/logging/docs/agent/logging/configuration#special-fields
		trace := fmt.Sprintf("projects/%s/traces/%s", h.projectID, traceID)
		rec.Add("logging.googleapis.com/trace", slog.StringValue(trace))
		if spanID != "" {
			rec.Add("logging.googleapis.com/spanId", slog.StringValue(spanID))
		}
	}
	return h.handler.Handle(ctx, rec)
}

func (h *CloudLoggingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CloudLoggingHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *CloudLoggingHandler) WithGroup(name string) slog.Handler {
	return &CloudLoggingHandler{handler: h.handler.WithGroup(name)}
}
