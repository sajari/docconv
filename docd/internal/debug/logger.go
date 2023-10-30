package debug

import (
	"context"
	"log/slog"
)

const DebugHeader = "X-Debug"

type debugHandler struct {
	slog.Handler
}

func NewDebugHandler(h slog.Handler) *debugHandler {
	return &debugHandler{Handler: h}
}

var _ slog.Handler = (*debugHandler)(nil)

func (h *debugHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if debugEnabledInContext(ctx) {
		return true
	}
	return h.Handler.Enabled(ctx, level)
}

func (h *debugHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &debugHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *debugHandler) WithGroup(name string) slog.Handler {
	return &debugHandler{Handler: h.Handler.WithGroup(name)}
}
