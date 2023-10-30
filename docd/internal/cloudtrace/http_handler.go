package cloudtrace

import "net/http"

type HTTPHandler struct {
	// Handler to wrap.
	Handler http.Handler
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := contextWithTraceInfo(r.Context(), r.Header.Get(CloudTraceContextHeader))

	h.Handler.ServeHTTP(w, r.WithContext(ctx))
}
