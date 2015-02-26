// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "os"
    "github.com/bjarneh/latinx"
)

func someSortOfHelp() {

    switch {
    case help:
        helpAndExit()
    case version:
        versionAndExit()
    case use:
        usage()
        os.Exit(0)
    case list:
        listAndExit()
    }

}

func versionAndExit() {
    fmt.Println("giconv 0.1")
    os.Exit(0)
}

func listAndExit() {
    for _, v := range latinx.Available() {
        fmt.Printf("%s\n", v)
    }
    os.Exit(0)
}

func usage() {

    var msg string = `
  General synopsis:

  giconv -f=<ENCODING> -t=<ENCODING> filename


  // Write ISO-8859-1 of input.txt in UTF-8 to stdout

  giconv -f=utf-8 -t=iso-8859-1 input.txt

  // Write UTF-8 of input.txt in Latin9/ISO-8859-15 to stdout

  giconv -f=iso-8859-15 -t=utf-8 input.txt

  // Write UTF-8 of input.txt in Latin9/ISO-8859-15 to output.txt

  giconv -f=iso-8859-15 -t=utf-8 < input.txt > output.txt
    `
    fmt.Printf("%s\n", msg)
}

func helpAndExit() {
    var msg string = `
  giconv - simple iconv clone for ISO 8859

  usage: giconv -f=<ENCODING> -t=<ENCODING> filename

  options:

    -h --help      : print this message and exit
    -v --version   : print version and exit
    -f --from-code : input encoding  (default: ISO-8859-15)
    -t --to-code   : output encoding (default: UTF-8)
    -l --list      : list available encodings
    -u --usage     : print typical usage of giconv
    -o --output    : write output to file (not stdout)
    ̀`
    fmt.Printf("%s\n", msg)

    os.Exit(0)
}
