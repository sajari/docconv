# docconv

[![Go reference](https://pkg.go.dev/badge/code.sajari.com/docconv/v2.svg)](https://pkg.go.dev/code.sajari.com/docconv/v2)
[![Build status](https://github.com/sajari/docconv/workflows/Go/badge.svg?branch=master)](https://github.com/sajari/docconv/actions)
[![Report card](https://goreportcard.com/badge/code.sajari.com/docconv/v2)](https://goreportcard.com/report/code.sajari.com/docconv/v2)
[![Sourcegraph](https://sourcegraph.com/github.com/sajari/docconv/v2/-/badge.svg)](https://sourcegraph.com/github.com/sajari/docconv/v2)

A Go wrapper library to convert PDF, DOC, DOCX, XML, HTML, RTF, ODT, Pages documents and images (see optional dependencies below) to plain text.

## Installation

If you haven't setup Go before, you first need to [install Go](https://golang.org/doc/install).

To fetch and build the code:

```console
$ go install code.sajari.com/docconv/v2/docd@latest
```

See `go help install` for details on the installation location of the installed `docd` executable. Make sure that the full path to the executable is in your `PATH` environment variable.

## Dependencies

- tidy
- wv
- popplerutils
- unrtf
- https://github.com/JalfResi/justext

### Debian-based Linux

```console
$ sudo apt-get install poppler-utils wv unrtf tidy
$ go get github.com/JalfResi/justext
```

### macOS

```console
$ brew install poppler-qt5 wv unrtf tidy-html5
$ go get github.com/JalfResi/justext
```

### Optional dependencies

To add image support to the `docconv` library you first need to [install and build gosseract](https://github.com/otiai10/gosseract/tree/v2.2.4).

Now you can add `-tags ocr` to any `go` command when building/fetching/testing `docconv` to include support for processing images:

```console
$ go get -tags ocr code.sajari.com/docconv/v2/...
```

This may complain on macOS, which you can fix by installing [tesseract](https://tesseract-ocr.github.io) via brew:

```console
$ brew install tesseract
```

## docd tool

The `docd` tool runs as either:

1.  a service on port 8888 (by default)

    Documents can be sent as a multipart POST request and the plain text (body) and meta information are then returned as a JSON object.

2.  a service exposed from within a Docker container

    This also runs as a service, but from within a Docker container.
    Official images are published at https://hub.docker.com/r/sajari/docd.

    Optionally you can build it yourself:

    ```console
    $ cd docd
    $ docker build -t docd .
    ```

3.  via the command line.

    Documents can be sent as an argument, e.g.

    ```console
    $ docd -input document.pdf
    ```

### Optional flags

- `addr` - the bind address for the HTTP server, default is ":8888"
- `readability-length-low` - sets the readability length low if the ?readability=1 parameter is set
- `readability-length-high` - sets the readability length high if the ?readability=1 parameter is set
- `readability-stopwords-low` - sets the readability stopwords low if the ?readability=1 parameter is set
- `readability-stopwords-high` - sets the readability stopwords high if the ?readability=1 parameter is set
- `readability-max-link-density` - sets the readability max link density if the ?readability=1 parameter is set
- `readability-max-heading-distance` - sets the readability max heading distance if the ?readability=1 parameter is set
- `readability-use-classes` - comma separated list of readability classes to use if the ?readability=1 parameter is set

### How to start the service

```console
$ # This runs on port 8000
$ docd -addr :8000
```

## Example usage (code)

Some basic code is shown below, but normally you would accept the file by HTTP or open it from the file system.

This should be enough to get you started though.

### Use case 1: run locally

> Note: this assumes you have the [dependencies](#dependencies) installed.

```go
package main

import (
	"fmt"

	"code.sajari.com/docconv/v2"
)

func main() {
	res, err := docconv.ConvertPath("your-file.pdf")
	if err != nil {
		// TODO: handle
	}
	fmt.Println(res)
}
```

### Use case 2: request over the network

```go
package main

import (
	"fmt"

	"code.sajari.com/docconv/v2/client"
)

func main() {
	// Create a new client, using the default endpoint (localhost:8888)
	c := client.New()

	res, err := client.ConvertPath(c, "your-file.pdf")
	if err != nil {
		// TODO: handle
	}
	fmt.Println(res)
}
```

Alternatively, via a `curl`:

```console
$ curl -s -F input=@your-file.pdf http://localhost:8888/convert
```
