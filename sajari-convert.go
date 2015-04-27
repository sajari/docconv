package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"io"
	"os"
	"regexp"
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
	reExtension, _ := regexp.Compile(".([a-z]+)$")
	if matches := reExtension.FindAllStringSubmatch(filename, -1); len(matches) > 0 && len(matches[0]) > 1 {
		switch matches[0][1] {
			case "doc":
				return "application/msword"
			case "docx":
				return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
			case "pages":
				return "application/vnd.apple.pages"
			case "pdf":
				return "application/pdf"
			case "rtf":
				return "application/rtf"
			case "xml":
				return "text/xml"
			case "xhtml":
			case "html":
				return "text/html"
			case "txt":
				return "text/plain"
		}
	}
	return "application/octet-stream"
}

// Convert a file to plain text & meta data
func convert(input io.Reader, mimeType string, readability bool) *Response {
	startClock := time.Now()
	response := new(Response)
	switch mimeType {
		case "application/msword", "application/vnd.ms-word":
			response.Body, response.Meta = ConvertDoc(input)

		case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
			response.Body, response.Meta = ConvertDocx(input)

		case "application/vnd.apple.pages", "application/x-iwork-pages-sffpages":
			response.Body, response.Meta = ConvertPages(input)

		case "application/pdf":
			response.Body, response.Meta = ConvertPdf(input)

		case "application/rtf", "application/x-rtf", "text/rtf", "text/richtext":
			response.Body, response.Meta = ConvertRtf(input)

		case "text/html":
			response.Body, response.Meta = ConvertHtml(input, readability)

		case "text/url":
			response.Body, response.Meta = ConvertUrl(input, readability)

		case "text/xml", "application/xml":
			response.Body, response.Meta = ConvertXml(input)

		case "text/plain":
			body, _ := ioutil.ReadAll(input)
			response.Body = string(body)
	}
	response.MSecs = uint32(time.Since(startClock).Nanoseconds() / 1000000)
	return response
}

func main() {
	flag.Parse()
	
	if *inputPath != "" {
		fmt.Print(string(convertPath(*inputPath)))
	} else {
		serve()
	}
}

// Convert a file given a path
func convertPath(path string) []byte {
	mimeType := mimeTypeByExtension(path)
	if *logLevel >= 1 {
		log.Println("Converting file: " + path + " (" + mimeType + ")")
	}
	
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open file: ", err)
	}
	defer file.Close()
	
	jsonStr, _ := json.Marshal(convert(file, mimeType, false))
	if *logLevel >= 2 {
		log.Println(string(jsonStr))
	}
	return jsonStr
}

// Start the conversion web service
func serve() {
	http.HandleFunc("/convert", func(writer http.ResponseWriter, request *http.Request) {

		// Get uploaded file
		file, info, err := request.FormFile("input")
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			log.Println("File upload", err)
			return
		}

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
			log.Println("Recieved file: " + info.Filename + " (" + mimeType + ")")
		}

		// Readability flag. Currently only used for HTML
		var readability bool = false
		if request.FormValue("readability") == "1" {
			readability = true
			if *logLevel >= 2 {
				log.Println("Readability is on")
			}
		}

		jsonStr, _ := json.Marshal(convert(file, mimeType, readability))
		if *logLevel >= 2 {
			log.Println(string(jsonStr))
		}
		fmt.Fprintf(writer, "%s", jsonStr)
	})

	// Start webserver
	log.Println("Setting log level to", *logLevel)
	log.Println("Starting Sajari convert on", *listenAddr)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
