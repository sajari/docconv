package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/errorreporting"

	"github.com/gorilla/mux"

	"code.sajari.com/docconv/v2"
	"code.sajari.com/docconv/v2/docd/internal"
	"code.sajari.com/docconv/v2/docd/internal/cloudtrace"
	"code.sajari.com/docconv/v2/docd/internal/debug"
)

var (
	listenAddr = flag.String("addr", ":8888", "The address to listen on (e.g. 127.0.0.1:8888)")

	inputPath = flag.String("input", "", "The file path to convert and exit; no server")

	jsonCloudLogging = flag.Bool("json-cloud-logging", false, "Whether or not to enable JSON Cloud Logging")

	cloudTraceGCPProjectID = flag.String("cloud-trace-gcp-project-id", "", "The GCP project to use for Cloud Trace")

	errorReporting                 = flag.Bool("error-reporting", false, "Whether or not to enable GCP Error Reporting")
	errorReportingGCPProjectID     = flag.String("error-reporting-gcp-project-id", "", "The GCP project to use for Error Reporting")
	errorReportingAppEngineService = flag.String("error-reporting-app-engine-service", "", "The App Engine service to use for Error Reporting")

	readabilityLengthLow          = flag.Int("readability-length-low", 70, "Sets the readability length low")
	readabilityLengthHigh         = flag.Int("readability-length-high", 200, "Sets the readability length high")
	readabilityStopwordsLow       = flag.Float64("readability-stopwords-low", 0.2, "Sets the readability stopwords low")
	readabilityStopwordsHigh      = flag.Float64("readability-stopwords-high", 0.3, "Sets the readability stopwords high")
	readabilityMaxLinkDensity     = flag.Float64("readability-max-link-density", 0.2, "Sets the readability max link density")
	readabilityMaxHeadingDistance = flag.Int("readability-max-heading-distance", 200, "Sets the readability max heading distance")
	readabilityUseClasses         = flag.String("readability-use-classes", "good,neargood", "Comma separated list of readability classes to use")
)

func main() {
	flag.Parse()

	l := slog.New(debug.NewDebugHandler(slog.Default().Handler()))

	if *jsonCloudLogging {
		gcpProjectID := *cloudTraceGCPProjectID
		if gcpProjectID == "" {
			gcpProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
		}
		if gcpProjectID == "" {
			l.Debug("GOOGLE_CLOUD_PROJECT env var not provided, looking up internal metadata service")
			var err error
			gcpProjectID, err = metadata.ProjectID()
			if err != nil {
				l.Error("Could not autodetect GCP project ID", "error", err)
				os.Exit(1)
			}
		}

		l = slog.New(debug.NewDebugHandler(cloudtrace.NewCloudLoggingHandler(gcpProjectID, slog.LevelInfo)))
		slog.SetDefault(l)
		l.Info("Cloud Trace GCP project ID", "projectID", gcpProjectID)
	}

	var er internal.ErrorReporter = &internal.NopErrorReporter{}
	if *errorReporting {
		if *errorReportingGCPProjectID == "" {
			*errorReportingGCPProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
		}
		if *errorReportingAppEngineService == "" {
			*errorReportingAppEngineService = os.Getenv("GAE_SERVICE")
		}
		var err error
		er, err = errorreporting.NewClient(context.Background(), *errorReportingGCPProjectID, errorreporting.Config{
			ServiceName: *errorReportingAppEngineService,
			OnError: func(err error) {
				l.Error("Could not report error to Error Reporting service", "error", err)
			},
		})
		if err != nil {
			l.Error("Could not create Error Reporting client", "error", err)
			os.Exit(1)
		}
	}

	cs := &convertServer{
		l:  l,
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
			l.Error("Could not convert file", "error", err, "path", *inputPath)
			os.Exit(1)
		}
		fmt.Print(string(resp.Body))
		return
	}

	serve(l, er, cs)
}

// Start the conversion web service
func serve(l *slog.Logger, er internal.ErrorReporter, cs *convertServer) {
	r := mux.NewRouter()
	r.HandleFunc("/convert", cs.convert)

	h := internal.RecoveryHandler(l, er)(r)
	h = &debug.HTTPHandler{Handler: h}
	h = &cloudtrace.HTTPHandler{Handler: h}

	l.Info("Go version " + runtime.Version())

	l.Info(fmt.Sprintf("HTTP server listening on %q...", *listenAddr))
	if err := http.ListenAndServe(*listenAddr, h); err != nil {
		l.Error("HTTP server ListenAndServe", "error", err)
	}
}
