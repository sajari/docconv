package cloudtrace

import "strings"

// The header specification is:
// "X-Cloud-Trace-Context: TRACE_ID/SPAN_ID;o=TRACE_TRUE"
const CloudTraceContextHeader = "X-Cloud-Trace-Context"

func parseHeader(value string) (traceID, spanID string) {
	var found bool
	traceID, after, found := strings.Cut(value, "/")
	if found {
		spanID, _, _ = strings.Cut(after, ";")
	}
	return
}
