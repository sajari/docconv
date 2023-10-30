package debug

import (
	"context"
	"net/http"
	"strconv"
)

type HTTPHandler struct {
	// Handler to wrap.
	Handler http.Handler
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if ok, _ := strconv.ParseBool(r.Header.Get(DebugHeader)); ok {
		ctx = context.WithValue(ctx, debugEnabledKey{}, true)
	}

	h.Handler.ServeHTTP(w, r.WithContext(ctx))
}
