sajari-convert
==============

A Golang wrapper library to convert PDF, DOC, DOCX, XML, HTML, RTF, etc to plain text

The compiled binary runs as a service on port 8888 by default. Documents can be sent as a multipart POST request and the plain text (body) and meta information are then returned as a JSON object

### Dependencies
tidy, wv, popplerutils, unrtf, github.com/JalfResi/justext

Example install of dependencies (not all systems):

```sudo apt-get install poppler-utils wv unrtf tidy``` 
```go get github.com/JalfResi/justext``` 

### Optional Flags
 - "addr" - the port to listen on, default is "8888"
 - "log-level" - can be 0, 1 or 2. Levels described below
 - "readability-length-low" - Sets the readability length low if the ?readability=1 parameter is set
 - "readability-length-high" - Sets the readability length high if the ?readability=1 parameter is set
 - "readability-stopwords-low" - Sets the readability stopwords low if the ?readability=1 parameter is set
 - "readability-stopwords-high" - Sets the readability stopwords high if the ?readability=1 parameter is set
 - "readability-max-link-density" - Sets the readability max link density if the ?readability=1 parameter is set
 - "readability-max-heading-distance" - Sets the readability max heading distance if the ?readability=1 parameter is set
 - "readability-use-classes - Comma separated list of readability classes to use if the ?readability=1 parameter is set

##### Log levels
 - "0" - will only log errors & critical info
 - "1" - will log each request as well
 - "2" - will also log the response payloads

### How to start the service
```./sajari-convert -log-level=0   # will only log errors & critical info ```

```./sajari-convert -addr=8000 -log-level=1   # will run on port 8000 and log each request as well ```

## FAQ

### How to install?
Compile the binary. Check the binary is executable and then launch as per above with relevant flag settings.

### Why run as a service?
You don't have to. However, if you have a large memory footprint (tens of GB) you may have trouble with forking caused by underlying Go functions when accessing the above mentioned dependencies. 

### How to ensure the service stays running?
This is an operating system question, but we use "upstart"

### Any sample code on how to call the service?
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