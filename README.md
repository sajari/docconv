# docconv
[![Build Status](https://travis-ci.org/nuveo/docconv.svg?branch=master)](https://travis-ci.org/nuveo/docconv)

A Go wrapper library to convert PDF, DOC, DOCX, XML, HTML, RTF, ODT, Pages documents and images (see optional dependencies below) to plain text.

## Installation

If you haven't setup Go before, you need to first set a `GOPATH` (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)).

To fetch and build the code:

    $ go get github.com/nuveo/docconv/...

This will also build the command line tool `docd` into `$GOPATH/bin` (assumed to be in your `PATH` already).

## Dependencies
tidy, wv, popplerutils, unrtf, https://github.com/JalfResi/justext

Example install of dependencies (not all systems):

    $ sudo apt-get install poppler-utils wv unrtf tidy
    $ go get github.com/JalfResi/justext

### Optional Dependencies

To add image support to the `docconv` library you first need to install and build https://github.com/otiai10/gosseract.  Now you can add `-tags ocr` to any `go` command when building/fetching `docconv` to include support for processing images:

    $ go get -tags ocr github.com/nuveo/docconv/...

## docd tool

The `docd` tool runs as either

1. a service on port 8888 (by default)

   Documents can be sent as a multipart POST request and the plain text (body) and meta information are then returned as a JSON object

2. via the command line.

   Documents can be sent as an argument, e.g.

   ```docd -input document.pdf```

### Optional Flags
 - "addr" - the bind address for the HTTP server, default is ":8888"
 - "log-level"
    - 0: errors & critical info
    - 1: inclues 0 and logs each request as well
    - 2: include 1 and logs the response payloads
 - "readability-length-low" - Sets the readability length low if the ?readability=1 parameter is set
 - "readability-length-high" - Sets the readability length high if the ?readability=1 parameter is set
 - "readability-stopwords-low" - Sets the readability stopwords low if the ?readability=1 parameter is set
 - "readability-stopwords-high" - Sets the readability stopwords high if the ?readability=1 parameter is set
 - "readability-max-link-density" - Sets the readability max link density if the ?readability=1 parameter is set
 - "readability-max-heading-distance" - Sets the readability max heading distance if the ?readability=1 parameter is set
 - "readability-use-classes - Comma separated list of readability classes to use if the ?readability=1 parameter is set

### How to start the service
```docd -log-level 0   # will only log errors & critical info ```

```docd -addr 8000 -log-level 1   # will run on port 8000 and log each request as well ```

## Example Usage (code)
Some basic code is shown below, but normally you would accept the file by http or open it from the file system. It should be enough to get you started though...

```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"bytes"
	"net/http"
	"mime/multipart"
	"net/textproto"
	"fmt"
)

type ConversionResponse struct {
	Body string             `json:"body"`
	Meta map[string]string  `json:"meta"`
	MSecs uint32            `json:"msecs"`
}

// Use the conversion service to convert data
func ConvertData(input []byte, mimeType string) ([]byte, map[string]string, error) {

	convertUrl := "http://localhost:8888/convert"
	convertParam := "input"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="`+convertParam+`"; filename="noname"`)
	h.Set("Content-Type", mimeType)
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, nil, err
	}
	_, err = part.Write(input)
	if err != nil {
		return nil, nil, err
	}
	err = writer.Close()
	if err != nil {
	  return nil, nil, err
	}
	client := &http.Client{}

	request, err := http.NewRequest("POST", convertUrl, body)
	if err != nil {
		return nil, nil, err
	}
	request.Header["Content-Type"] = []string{"multipart/form-data; boundary="+writer.Boundary()}
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	jsonBlob, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	converted := new(ConversionResponse)
	err = json.Unmarshal(jsonBlob, &converted)
	if err != nil {
		return nil, nil, err
	}
	return []byte(converted.Body), converted.Meta, nil
}

func main() {
	input := []byte{} // This would be the file contents
	mimeType := "application/pdf" // Also pass the mimetype of the file
	body, meta, err := ConvertData(input, mimeType)
	fmt.Println("The body text is : ", body)
	fmt.Println("The file meta data is a map : ", meta)
	fmt.Println("Any errors are returned here : ", err)
}
```
