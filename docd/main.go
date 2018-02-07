package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"code.sajari.com/docconv"
)

var (
	inputPath                     = flag.String("input", "", "The file path to convert and exit; no server")
	listenAddr                    = flag.String("addr", ":8888", "The address to listen on (e.g. 127.0.0.1:8888)")
	logLevel                      = flag.Uint("log-level", 0, "The verbosity of the log")
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
			log.Printf("error converting file '%v': %v", *inputPath, err)
		}
		fmt.Print(string(resp.Body))
		return
	}
	serve()
}

// Convert a file given a path
func convertPath(path string, readability bool) ([]byte, error) {
	mimeType := docconv.MimeTypeByExtension(path)
	if *logLevel >= 1 {
		log.Println("Converting file: " + path + " (" + mimeType + ")")
	}

	// Open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := docconv.Convert(f, mimeType, readability)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if *logLevel >= 2 {
		log.Println(string(b))
	}
	return b, nil
}

// Start the conversion web service
func serve() {
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		// Readability flag. Currently only used for HTML
		var readability bool
		if r.FormValue("readability") == "1" {
			readability = true
			if *logLevel >= 2 {
				log.Println("Readability is on")
			}
		}

		path := r.FormValue("path")
		if path != "" {
			b, err := docconv.ConvertPathReadability(path, readability)
			if err != nil {
				// TODO: return a sensible status code for errors like this.
				log.Printf("error converting path '%v': %v", path, err)
				return
			}
			w.Write(b)
			return
		}

		// Get uploaded file
		file, info, err := r.FormFile("input")
		if err != nil {
			log.Println("File upload", err)
			return
		}
		defer file.Close()

		// Abort if file doesn't have a mime type
		if len(info.Header["Content-Type"]) == 0 {
			log.Println("No content type", info.Filename)
			return
		}

		// If a generic mime type was provided then use file extension to determine mimetype
		mimeType := info.Header["Content-Type"][0]
		if mimeType == "application/octet-stream" {
			mimeType = docconv.MimeTypeByExtension(info.Filename)
		}

		if *logLevel >= 1 {
			log.Println("Received file: " + info.Filename + " (" + mimeType + ")")
		}

		data, err := docconv.Convert(file, mimeType, readability)
		if err != nil {
			log.Printf("error converting data: %v", err)
			data = &docconv.Response{
				Error: err.Error(),
			}
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("error marshaling JSON data: %v", err)
			return
		}
	})

	// Start webserver
	log.Println("Setting log level to", *logLevel)
	log.Println("Starting docconv on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
