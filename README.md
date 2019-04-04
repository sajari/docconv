# docconv

## Notes about this fork

This is a fork of https://github.com/sajari/docconv . It removes support for HTML conversion because the original project
relies on https://github.com/JalfResi/justext which has [licensing issues](https://github.com/JalfResi/justext/issues/32).

## Introduction

A Go wrapper library to convert PDF, DOC, DOCX, XML, RTF, ODT, Pages documents and images (see optional dependencies below) to plain text.

## Installation

If you haven't setup Go before, you need to first set a `GOPATH` (see [https://golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH)).

To fetch and build the code:

    $ go get code.sajari.com/docconv/...

This will also build the command line tool `docd` into `$GOPATH/bin` (assumed to be in your `PATH` already).

## Dependencies
tidy, wv, popplerutils, unrtf, https://github.com/JalfResi/justext

Example install of dependencies (not all systems):

    $ sudo apt-get install poppler-utils wv unrtf tidy
    $ go get github.com/JalfResi/justext

### Optional Dependencies

To add image support to the `docconv` library you first need to install and build https://github.com/otiai10/gosseract.  Now you can add `-tags ocr` to any `go` command when building/fetching `docconv` to include support for processing images:

    $ go get -tags ocr code.sajari.com/docconv/...

This may complain on OSX, which you can fix by installing tesseract via brew:
	
	$ brew install tesseract

## docd tool

The `docd` tool runs as either

1. a service on port 8888 (by default)

   Documents can be sent as a multipart POST request and the plain text (body) and meta information are then returned as a JSON object

2. a service exposed from within a Docker container

   This also runs as a service, but from within a Docker container. There are three build scripts: 
   [./docd/debian.sh](./docd/debian.sh)
   [./docd/alpine.sh](./docd/alpine.sh)
   [./docd/appengine.sh](./docd/appengine.sh)

   The debian version uses the Debian package repository which can vary with builds. The Alpine version uses a very cut down Linux distribution to produce a container ~40MB. It also locks the dependency versions for consistency, but may miss out on future updates. The appengine version is a flex based custom runtime for Google Cloud.

3. via the command line.

   Documents can be sent as an argument, e.g.

   ```docd -input document.pdf```

### Optional Flags
 - "addr" - the bind address for the HTTP server, default is ":8888"
 - "log-level"
    - 0: errors & critical info
    - 1: inclues 0 and logs each request as well
    - 2: include 1 and logs the response payloads

### How to start the service
```docd -log-level 0   # will only log errors & critical info ```

```docd -addr :8000 -log-level 1   # will run on port 8000 and log each request as well ```

## Example Usage (code)
Some basic code is shown below, but normally you would accept the file by http or open it from the file system. It should be enough to get you started though...

Use case 1: run locally 
Note: this assumes you have the dependencies installed.

```go
package main

import (
	"fmt"
	"log"

	"github.com/monzo/docconv"
)

func main() {
	res, err := docconv.ConvertPath("your-file.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
```

Use case 2: request over the network

```go
package main

import (
	"fmt"
	"log"

	"github.com/monzo/docconv/client"
)

func main() {
	// Create a new client, using the default endpoint (localhost:8888)
	c := client.New()

	res, err := client.ConvertPath(c, "your-file.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
```
