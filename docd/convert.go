package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"cloud.google.com/go/errorreporting"

	"code.sajari.com/docconv"
	"code.sajari.com/docconv/docd/internal"
)

type convertServer struct {
	l  *slog.Logger
	er internal.ErrorReporter
}

func (s *convertServer) convert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Readability flag. Currently only used for HTML
	var readability bool
	if r.FormValue("readability") == "1" {
		readability = true
		s.l.DebugContext(ctx, "Readability is on")
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

	s.l.InfoContext(ctx, "Received file", "filename", info.Filename, "mimeType", mimeType)

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

	s.l.InfoContext(ctx, fmt.Sprintf(pattern, args...))
}

func (s *convertServer) serverError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"error":"internal server error"}`))

	e := errorreporting.Entry{
		Error: err,
		Req:   r,
	}
	s.er.Report(e)

	s.l.ErrorContext(ctx, err.Error(), "error", err)
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
