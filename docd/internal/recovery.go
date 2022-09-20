package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"cloud.google.com/go/errorreporting"
)

type recoveryHandler struct {
	er      ErrorReporter
	handler http.Handler
}

// RecoveryHandler is HTTP middleware that recovers from a panic, writes a
// 500, reports the panic, logs the panic and continues to the next handler.
func RecoveryHandler(er ErrorReporter) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &recoveryHandler{er: er, handler: h}
	}
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"internal server error"}`))
			h.handle(req, &recovered{rec, debug.Stack()})
		}
	}()

	h.handler.ServeHTTP(w, req)
}

func (h recoveryHandler) handle(r *http.Request, err error) {
	stack, _ := stackFromRecovered(err)

	e := errorreporting.Entry{
		Error: err,
		Stack: stack,
		Req:   r,
	}
	h.er.Report(e)

	log.Println(err)
	log.Printf("%s", stack)
}

// recovered represents the return value from a call to recover.
type recovered struct {
	// p is the error value passed to the call of panic.
	p interface{}
	// stack is the panic stack trace.
	stack []byte
}

var _ error = (*recovered)(nil)

// Error implements error.
func (e *recovered) Error() string {
	if err, ok := e.p.(error); ok {
		return err.Error()
	}
	return fmt.Sprintf("panic: %v", e.p)
}

// stackFromRecovered returns a stack trace and true if the recovdered has a
// stack trace created by this package.
//
// Otherwise it returns nil and false.
func stackFromRecovered(err error) ([]byte, bool) {
	var rec *recovered
	if errors.As(err, &rec) {
		return rec.stack, true
	}
	return nil, false
}
