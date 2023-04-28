package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"cloud.google.com/go/errorreporting"
	"code.sajari.com/docconv"

	"github.com/gorilla/mux"
)

var (
	listenAddr = flag.String("addr", ":8888", "The address to listen on (e.g. 127.0.0.1:8888)")

	inputPath = flag.String("input", "", "The file path to convert and exit; no server")

	readabilityLengthLow          = flag.Int("readability-length-low", 70, "Sets the readability length low")
	readabilityLengthHigh         = flag.Int("readability-length-high", 200, "Sets the readability length high")
	readabilityStopwordsLow       = flag.Float64("readability-stopwords-low", 0.2, "Sets the readability stopwords low")
	readabilityStopwordsHigh      = flag.Float64("readability-stopwords-high", 0.3, "Sets the readability stopwords high")
	readabilityMaxLinkDensity     = flag.Float64("readability-max-link-density", 0.2, "Sets the readability max link density")
	readabilityMaxHeadingDistance = flag.Int("readability-max-heading-distance", 200, "Sets the readability max heading distance")
	readabilityUseClasses         = flag.String("readability-use-classes", "good,neargood", "Comma separated list of readability classes to use")
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

type convertServer struct {
	er ErrorReporter
}

func (s *convertServer) ping(w http.ResponseWriter, r *http.Request) {
	s.respond(r.Context(), w, r, http.StatusOK, "pong")
}

func (s *convertServer) convert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Readability flag. Currently only used for HTML
	var readability bool
	if r.FormValue("readability") == "1" {
		readability = true
		log.Println("Readability is on")

	}

	path := r.FormValue("path")
	if path != "" {
		mimeType := docconv.MimeTypeByExtension(path)

		f, err := os.Open(path)
		if err != nil {
			s.serverError(ctx, w, r, fmt.Errorf("could not open file: %w", err))
			return
		}
		defer f.Close()

		data, err := docconv.Convert(f, mimeType, readability)
		if err != nil {
			s.serverError(ctx, w, r, fmt.Errorf("could not convert file from path %v: %w", path, err))
			return
		}

		s.respond(ctx, w, r, http.StatusOK, data)
		return
	}

	// Get uploaded file
	file, info, err := r.FormFile("input")
	if err != nil {
		s.serverError(ctx, w, r, fmt.Errorf("could not get input file: %w", err))
		return
	}
	defer file.Close()

	// Abort if file doesn't have a mime type
	if len(info.Header["Content-Type"]) == 0 {
		s.clientError(ctx, w, r, http.StatusUnprocessableEntity, "input file %v does not have a Content-Type header", info.Filename)
		return
	}

	// If a generic mime type was provided then use file extension to determine mimetype
	mimeType := info.Header["Content-Type"][0]
	if mimeType == "application/octet-stream" {
		mimeType = docconv.MimeTypeByExtension(info.Filename)
	}

	log.Printf("Received file: %v (%v)", info.Filename, mimeType)

	data, err := docconv.Convert(file, mimeType, readability)
	if err != nil {
		s.serverError(ctx, w, r, fmt.Errorf("could not convert file: %w", err))
		return
	}

	s.respond(ctx, w, r, http.StatusOK, data)
}

func (s *convertServer) clientError(ctx context.Context, w http.ResponseWriter, r *http.Request, code int, pattern string, args ...interface{}) {
	s.respond(ctx, w, r, code, &docconv.Response{
		Error: fmt.Sprintf(pattern, args...),
	})

	log.Printf(pattern, args...)
}

func (s *convertServer) serverError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error":"internal server error"}`))

	e := errorreporting.Entry{
		Error: err,
		Req:   r,
	}
	s.er.Report(e)

	log.Printf("%v", err)
}

func (s *convertServer) respond(ctx context.Context, w http.ResponseWriter, r *http.Request, code int, resp interface{}) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(resp)
	if err != nil {
		s.serverError(ctx, w, r, fmt.Errorf("could not marshal JSON response: %w", err))
		return
	}
	w.WriteHeader(code)
	n, err := io.Copy(w, buf)
	if err != nil {
		panic(fmt.Errorf("could not write to response (failed after %d bytes): %w", n, err))
	}
}

func main() {
	flag.Parse()

	var er ErrorReporter = &NopErrorReporter{}

	cs := &convertServer{
		er: er,
	}

	// TODO: Improve this (remove the need for it!)
	docconv.HTMLReadabilityOptionsValues = docconv.HTMLReadabilityOptions{
		LengthLow:             *readabilityLengthLow,
		LengthHigh:            *readabilityLengthHigh,
		StopwordsLow:          *readabilityStopwordsLow,
		StopwordsHigh:         *readabilityStopwordsHigh,
		MaxLinkDensity:        *readabilityMaxLinkDensity,
		MaxHeadingDistance:    *readabilityMaxHeadingDistance,
		ReadabilityUseClasses: *readabilityUseClasses,
	}

	if *inputPath != "" {
		resp, err := docconv.ConvertPath(*inputPath)
		if err != nil {
			log.Println("Could not convert file", "error", err, "path", *inputPath)
			os.Exit(1)
		}
		fmt.Print(string(resp.Body))
		return
	}

	serve(er, cs)
}

// Start the conversion web service
func serve(er ErrorReporter, cs *convertServer) {
	r := mux.NewRouter()
	r.HandleFunc("/convert", cs.convert)
	r.HandleFunc("/ping", cs.ping)

	// Start webserver
	log.Println("Starting docconv on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, RecoveryHandler(er)(r)))
}
