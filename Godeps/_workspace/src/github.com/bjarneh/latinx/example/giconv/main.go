// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/bjarneh/latinx"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// "iso88591" => latinx.ISO_8859_1 and so on..
var encode map[string]int

// either os.Stdin or result of Open
var reader io.Reader

// either os.Stdout or result of Open
var writer io.Writer

var UTF8 int = -1 // distinguish UTF8 from ISO_8859_*

// boolean flags
var (
	help    = false
	version = false
	list    = false
	use     = false
)

// string flags with defaults
var (
	from   = "iso-8859-1"
	to     = "utf-8"
	output = ""
)

func init() {

	flag.StringVar(&to, "t", to, "output encoding")
	flag.StringVar(&to, "to", to, "output encoding") // shorter alias 
	flag.StringVar(&to, "to-code", to, "output encoding")
	flag.StringVar(&from, "f", from, "input encoding")
	flag.StringVar(&from, "from", from, "input encoding") // shorter alias
	flag.StringVar(&from, "from-code", from, "input encoding")
	flag.StringVar(&output, "o", output, "write result to output")
	flag.StringVar(&output, "output", output, "write results to output")
	flag.BoolVar(&help, "h", help, "print help")
	flag.BoolVar(&help, "help", help, "print help")
	flag.BoolVar(&version, "v", version, "print version")
	flag.BoolVar(&version, "version", version, "print version")
	flag.BoolVar(&list, "l", list, "list available encodings")
	flag.BoolVar(&list, "list", list, "list available encodings")
	flag.BoolVar(&use, "u", use, "short usage example")
	flag.BoolVar(&use, "usage", use, "short usage example")

	// override flag.Usage since it repeats flags in some sense :-)
	flag.Usage = usage

	// use this to fetch encoding from input string
	encode = make(map[string]int)
	encode["latin1"] = latinx.ISO_8859_1
	encode["iso88591"] = latinx.ISO_8859_1
	encode["latin2"] = latinx.ISO_8859_2
	encode["iso88592"] = latinx.ISO_8859_2
	encode["latin3"] = latinx.ISO_8859_3
	encode["iso88593"] = latinx.ISO_8859_3
	encode["latin4"] = latinx.ISO_8859_4
	encode["iso88594"] = latinx.ISO_8859_4
	encode["cyrillic"] = latinx.ISO_8859_5
	encode["iso88595"] = latinx.ISO_8859_5
	encode["arabic"] = latinx.ISO_8859_6
	encode["iso88596"] = latinx.ISO_8859_6
	encode["greek"] = latinx.ISO_8859_7
	encode["iso88597"] = latinx.ISO_8859_7
	encode["hebrew"] = latinx.ISO_8859_8
	encode["iso88598"] = latinx.ISO_8859_8
	encode["latin5"] = latinx.ISO_8859_9
	encode["iso88599"] = latinx.ISO_8859_9
	encode["latin6"] = latinx.ISO_8859_10
	encode["iso885910"] = latinx.ISO_8859_10
	encode["thai"] = latinx.ISO_8859_11
	encode["iso885911"] = latinx.ISO_8859_11
	encode["latin7"] = latinx.ISO_8859_13
	encode["iso885913"] = latinx.ISO_8859_13
	encode["latin8"] = latinx.ISO_8859_14
	encode["iso885914"] = latinx.ISO_8859_14
	encode["latin9"] = latinx.ISO_8859_15
	encode["iso885915"] = latinx.ISO_8859_15
	encode["latin10"] = latinx.ISO_8859_16
	encode["iso885916"] = latinx.ISO_8859_16
	encode["windows1252"] = latinx.Windows1252
	encode["cp11252"] = latinx.Windows1252
	// special case
	encode["utf8"] = UTF8
}

// this is where the interesting stuff happens..
func main() {

	var inputbytes, utf8bytes, outputbytes []byte

	flag.Parse()

	// print help/version/list/usage...
	someSortOfHelp()

	// read input from stdin or file
	inputbytes = getInputBytes()

	// convert inputformat to UTF-8 (decode)
	utf8bytes = getUtf8FromInput(inputbytes)

	// convert UTF-8 to output format (encode)
	outputbytes = convertToDesired(utf8bytes)

	// write output to stdout or file
	writeOutput(outputbytes)

}

func writeOutput(outputbytes []byte) {

	var file *os.File
	var err error

	if output == "" {

		writer = os.Stdout

	} else {

		file, err = os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		defer file.Close()
		writer = file

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	_, err = writer.Write(outputbytes)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func getUtf8FromInput(inputbytes []byte) (utf8bytes []byte) {

	var fromId int
	var converter *latinx.Converter
	var err error

	fromId = converterIdent(from)

	if fromId != UTF8 {

		converter = latinx.Get(fromId)
		utf8bytes, err = converter.Decode(inputbytes)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

	} else {
		utf8bytes = inputbytes
	}

	return
}

func convertToDesired(utf8bytes []byte) (outputbytes []byte) {

	var toId int
	var err error
	var converter *latinx.Converter

	toId = converterIdent(to)

	if toId != UTF8 {

		converter = latinx.Get(toId)
		outputbytes, _, err = converter.Encode(utf8bytes)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}

	} else {
		outputbytes = utf8bytes
	}

	return
}

func getInputBytes() (input []byte) {

	var err error
	var file *os.File

	if flag.NArg() == 0 {

		reader = os.Stdin

	} else {

		file, err = os.Open(flag.Arg(0))
		defer file.Close()
		reader = file

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}

	input, err = ioutil.ReadAll(reader)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	return
}

func converterIdent(encoding string) (iencoding int) {

	// remove '-','_',' ' and lowercase string
	lower := strings.ToLower(encoding)
	reg := regexp.MustCompile("[\\- _]")
	key := reg.ReplaceAllString(lower, "")

	iencoding, ok := encode[key]
	if !ok {
		fmt.Fprintf(os.Stderr, "%s - unknown encoding\n", encoding)
		os.Exit(1)
	}

	return iencoding
}
