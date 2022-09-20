package internal

import (
	"io"

	"cloud.google.com/go/errorreporting"
)

// ErrorReporter reports errors.
type ErrorReporter interface {
	Report(errorreporting.Entry)
	io.Closer
}

// NopErrorReporter is a no-op reporter.
type NopErrorReporter struct{}

var _ ErrorReporter = (*NopErrorReporter)(nil)

// Report implements ErrorReporter.
func (r *NopErrorReporter) Report(e errorreporting.Entry) {}

// Close implements ErrorReporter.
func (r *NopErrorReporter) Close() error { return nil }
