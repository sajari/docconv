package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

// Response payload sent back to the requestor
type Response struct {
	Body  string            `json:"body"`
	Meta  map[string]string `json:"meta"`
	MSecs uint32            `json:"msecs"`
}

var (
	inputPath                     *string  = flag.String("input", "", "The file path to convert and exit; no server")
	listenAddr                    *string  = flag.String("addr", ":8888", "The address to listen on (e.g. 127.0.0.1:8888)")
	logLevel                      *uint    = flag.Uint("log-level", 0, "The verbosity of the log")
	readabilityLengthLow          *int     = flag.Int("readability-length-low", 70, "Sets the readability length low")
	readabilityLengthHigh         *int     = flag.Int("readability-length-high", 200, "Sets the readability length high")
	readabilityStopwordsLow       *float64 = flag.Float64("readability-stopwords-low", 0.2, "Sets the readability stopwords low")
	readabilityStopwordsHigh      *float64 = flag.Float64("readability-stopwords-high", 0.3, "Sets the readability stopwords high")
	readabilityMaxLinkDensity     *float64 = flag.Float64("readability-max-link-density", 0.2, "Sets the readability max link density")
	readabilityMaxHeadingDistance *int     = flag.Int("readability-max-heading-distance", 200, "Sets the readability max heading distance")
	readabilityUseClasses         *string  = flag.String("readability-use-classes", "good,neargood", "Comma separated list of readability classes to use")
)

// Determine the mime type by the file's extension
func mimeTypeByExtension(filename string) string {
	switch path.Ext(filename) {
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".odt":
		return "application/vnd.oasis.opendocument.text"
	case ".pages":
		return "application/vnd.apple.pages"
	case ".pdf":
		return "application/pdf"
	case ".rtf":
		return "application/rtf"
	case ".xml":
		return "text/xml"
	case ".xhtml", ".html", ".htm":
		return "text/html"
	case ".txt":
		return "text/plain"
	}
	return "application/octet-stream"
}

// Convert a file to plain text & meta data
func convert(r io.Reader, mimeType string, readability bool) *Response {
	start := time.Now()

	var body string
	var meta map[string]string
	var err error
	switch mimeType {
	case "application/msword", "application/vnd.ms-word":
		body, meta = ConvertDoc(r)

	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		body, meta = ConvertDocx(r)

	case "application/vnd.oasis.opendocument.text":
		body, meta = ConvertODT(r)

	case "application/vnd.apple.pages", "application/x-iwork-pages-sffpages":
		body, meta, err = ConvertPages(r)

	case "application/pdf":
		body, meta = ConvertPDF(r)

	case "application/rtf", "application/x-rtf", "text/rtf", "text/richtext":
		body, meta, err = ConvertRTF(r)

	case "text/html":
		body, meta, err = ConvertHTML(r, readability)

	case "text/url":
		body, meta = ConvertURL(r, readability)

	case "text/xml", "application/xml":
		body, meta, err = ConvertXML(r)

	case "text/plain":
		// TODO: Don't ignore the error.
		b, _ := ioutil.ReadAll(r)
		body = string(b)
	}

	if err != nil {
		// TODO(dhowden): Don't log this, actually return it in the response.
		log.Printf("could not convert file: %v", err)
	}

	return &Response{
		Body:  body,
		Meta:  meta,
		MSecs: uint32(time.Since(start) / time.Millisecond),
	}
}

func main() {
	flag.Parse()

	if *inputPath != "" {
		b, err := convertPath(*inputPath, false)
		if err != nil {
			log.Fatal("Cannot open file: ", err)
		}
		fmt.Print(string(b))
		return
	}
	serve()
}

// Convert a file given a path
func convertPath(path string, readability bool) ([]byte, error) {
	mimeType := mimeTypeByExtension(path)
	if *logLevel >= 1 {
		log.Println("Converting file: " + path + " (" + mimeType + ")")
	}

	// Open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// TODO: Don't ignore this error.
	b, _ := json.Marshal(convert(f, mimeType, readability))
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
			b, err := convertPath(path, readability)
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
			mimeType = mimeTypeByExtension(info.Filename)
		}

		if *logLevel >= 1 {
			log.Println("Received file: " + info.Filename + " (" + mimeType + ")")
		}

		jsonStr, _ := json.Marshal(convert(file, mimeType, readability))
		if *logLevel >= 2 {
			log.Println(string(jsonStr))
		}
		fmt.Fprintf(w, "%s", jsonStr)
	})

	// Start webserver
	log.Println("Setting log level to", *logLevel)
	log.Println("Starting Sajari convert on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
