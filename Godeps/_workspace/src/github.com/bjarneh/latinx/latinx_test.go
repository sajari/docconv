// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx_test

import (
    "testing"
    "github.com/bjarneh/latinx"
    "path/filepath"
    "io/ioutil"
    "os"
)


// compare to byte arrays
func byteEqual(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i := 0; i < len(b); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

type tuple struct {
    filename  string
    converter *latinx.Converter
}

func TestBothWaysAll(t *testing.T) {

    var tuples []*tuple
    var fname, srcroot, me string

    srcroot = os.Getenv("SRCROOT")
    me = "github.com/bjarneh/latinx"

    tuples = make([]*tuple, 0)

    tuples = append(tuples, &tuple{"part1.go",
        latinx.Get(latinx.ISO_8859_1)})
    tuples = append(tuples, &tuple{"part2.go",
        latinx.Get(latinx.ISO_8859_2)})
    tuples = append(tuples, &tuple{"part3.go",
        latinx.Get(latinx.ISO_8859_3)})
    tuples = append(tuples, &tuple{"part4.go",
        latinx.Get(latinx.ISO_8859_4)})
    tuples = append(tuples, &tuple{"part5.go",
        latinx.Get(latinx.ISO_8859_5)})
    tuples = append(tuples, &tuple{"part6.go",
        latinx.Get(latinx.ISO_8859_6)})
    tuples = append(tuples, &tuple{"part7.go",
        latinx.Get(latinx.ISO_8859_7)})
    tuples = append(tuples, &tuple{"part8.go",
        latinx.Get(latinx.ISO_8859_8)})
    tuples = append(tuples, &tuple{"part9.go",
        latinx.Get(latinx.ISO_8859_9)})
    tuples = append(tuples, &tuple{"part10.go",
        latinx.Get(latinx.ISO_8859_10)})
    tuples = append(tuples, &tuple{"part11.go",
        latinx.Get(latinx.ISO_8859_11)})
    tuples = append(tuples, &tuple{"part13.go",
        latinx.Get(latinx.ISO_8859_13)})
    tuples = append(tuples, &tuple{"part14.go",
        latinx.Get(latinx.ISO_8859_14)})
    tuples = append(tuples, &tuple{"part15.go",
        latinx.Get(latinx.ISO_8859_15)})
    tuples = append(tuples, &tuple{"part16.go",
        latinx.Get(latinx.ISO_8859_16)})
    tuples = append(tuples, &tuple{"windows1252.go",
        latinx.Get(latinx.Windows1252)})

    for _, tup := range tuples {

        // testing with godag
        if srcroot != "" {
            fname = filepath.Join(srcroot, me, tup.filename)
        } else {
            fname = tup.filename
        }

        utf, err := ioutil.ReadFile(fname)

        if err != nil {
            t.Fatalf("%s\n", err)
        }

        lat, _, err := tup.converter.Encode(utf)

        if err != nil {
            t.Fatalf("%s\n", err)
        }

        u2, err := tup.converter.Decode(lat)

        if err != nil {
            t.Fatalf("%s\n", err)
        }

        if !byteEqual(u2, utf) {
            t.Fatalf("failed %s.encode().decode()|%s\n",
                tup.filename, tup.converter)
        }
    }
}

func TestPartial(t *testing.T) {

    conv := latinx.Get(latinx.ISO_8859_1)

    s := "øæå" // 6 bytes == 3 utf-8 multibytes
    sb := []byte(s)

    latin, size, err := conv.Encode(sb[0:3])

    if err == nil {
        t.Fatalf("failed: conv.Encode(partial) => err == nil\n")
    }

    if err != latinx.PARTIAL {
        t.Fatalf("failed: err != latinx.PARTIAL")
    }

    _, ok := err.(latinx.UnicodeError)

    if !ok {
        t.Fatalf("failed: os.Error is not UnicodeError\n")
    }

    if size != 2 {
        t.Fatalf("failed: conv.Encode(partial) => size is wrong\n")
    }

    if len(latin) != 1 {
        t.Fatalf("failed: conv.Encode(partial) => len(latin) is wrong\n")
    }

    if latin[0] != 0xF8 {
        t.Fatalf("failed: conv.Encode(partial) => symbol is wrong\n")
    }
}
