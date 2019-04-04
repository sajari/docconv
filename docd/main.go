package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/monzo/docconv"
	"log"
	"net/http"
)

var (
	inputPath                     = flag.String("input", "", "The file path to convert and exit; no server")
	listenAddr                    = flag.String("addr", ":8888", "The address to listen on (e.g. 127.0.0.1:8888)")
	logLevel                      = flag.Uint("log-level", 0, "The verbosity of the log")
)

func main() {
	flag.Parse()

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
