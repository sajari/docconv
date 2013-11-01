package main

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"flag"
	"io/ioutil"
	"encoding/json"
	"regexp"
)

// Response payload sent back to the requestor
type Response struct {
	Body string             `json:"body"`
	Meta map[string]string  `json:"meta"`
	MSecs uint32            `json:"msecs"`
}

func main() {
	var listenAddr string
	var logLevel uint
	flag.StringVar(&listenAddr, "addr", ":8888", "The address to listen on (e.g. 127.0.0.1:8888)")
	flag.UintVar(&logLevel, "log-level", 0, "The verbosity of the log")
	flag.Parse()

	// Accept requests at /convert
	http.HandleFunc("/convert", func(writer http.ResponseWriter, request *http.Request) {
		startClock := time.Now()

		// Get uploaded file
        file, info, err := request.FormFile("input")
        defer file.Close()
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
			reExtension, _ := regexp.Compile(".([a-z]+)$")
			if matches := reExtension.FindAllStringSubmatch(info.Filename, -1); len(matches) > 0 && len(matches[0]) > 1 {
				switch matches[0][1] {
					case "doc":
						mimeType = "application/msword"
					case "docx":
						mimeType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
					case "pdf":
						mimeType = "application/pdf"
					case "rtf":
						mimeType = "application/rtf"
					case "xml":
						mimeType = "text/xml"
					case "xhtml":
					case "html":
						mimeType = "text/html"
					case "txt":
						mimeType = "text/plain"
				}
			}
		}

		if logLevel >= 1 {
			log.Println("Recieved file: "+info.Filename+" ("+mimeType+")")
		}
		
		// Convert document
		response := new(Response)
		switch mimeType {
			case "application/msword", "application/vnd.ms-word":
				response.Body, response.Meta = ConvertDoc(file)

			case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
				response.Body, response.Meta = ConvertDocx(file)

			case "application/pdf":
				response.Body, response.Meta = ConvertPdf(file)

			case "application/rtf", "application/x-rtf", "text/rtf", "text/richtext":
				response.Body, response.Meta = ConvertRtf(file)

			case "text/html":
				response.Body, response.Meta = ConvertHtml(file)

			case "text/xml", "application/xml":
				response.Body, response.Meta = ConvertXml(file)

			case "text/plain":
				body, _ := ioutil.ReadAll(file)
				response.Body = string(body)
		}
		
		// Return result
		response.MSecs = uint32(time.Since(startClock).Nanoseconds() / 1000000)
		jsonStr, _ := json.Marshal(response)
		if logLevel >= 2 {
			log.Println(string(jsonStr))
		}
		fmt.Fprintf(writer, "%s", jsonStr)
	})

	// Start webserver
	log.Println("Setting log level to", logLevel)
	log.Println("Starting Sajari convert on", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
