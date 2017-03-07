// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

import (
	"bytes"
	"io"
)

// A LatinReader reads ISO-8859 streams from underlying reader, 
// decodes them to UTF-8, and writes them to a *bytes.Buffer, which
// is used to store the possibly larger byte-stream. After the decoded
// stream has been written to buffer, a Read(p []byte) from *bytes.Buffer
// is preformed.
type LatinReader struct {
	reader    io.Reader
	converter *Converter
	buf       *bytes.Buffer
}

// Initialize a new LatinReader using an underlying io.Reader and one of the
// available charsets (ISO_8895_1, ISO_8895_2..)
func NewReader(charset int, r io.Reader) io.Reader {
	conv, ok := converters[charset]
	if !ok {
		return nil
	}
	b := bytes.NewBuffer(make([]byte, 1024))
	b.Truncate(0)
	return &LatinReader{reader: r, converter: conv, buf: b}
}

// Read from underlying io.Reader and decode to UTF-8 and return result.
func (r *LatinReader) Read(p []byte) (n int, err error) {

	var p2, utf8bytes []byte
	var n2 int
	var e2, e3 error

	p2 = make([]byte, len(p))
	n2, e2 = r.reader.Read(p2)

	if e2 == nil || e2 == io.EOF {

		utf8bytes, e3 = r.converter.Decode(p2[:n2])

		if e3 != nil {
			return 0, e3
		}

		r.buf.Write(utf8bytes) // n always len(input), error always nil

		return r.buf.Read(p)

	} else {
		copy(p, p2)
		n = n2
		err = e2
	}

	return
}
