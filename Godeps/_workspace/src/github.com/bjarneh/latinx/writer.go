// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

import "io"

// A LatinWriter writer will encode UTF-8 byte-streams into selected
// ISO 8859 byte-stream, before writing them to underlying io.Writer.
type LatinWriter struct {
	writer    io.Writer
	converter *Converter
}

// Initialize a new LatinWriter with an underlying io.Writer and one
// of the available charsets (ISO_8895_1,ISO_8895_2...).
func NewWriter(charset int, w io.Writer) io.Writer {
	conv, ok := converters[charset]
	if !ok {
		return nil
	}
	return &LatinWriter{writer: w, converter: conv}
}

// The returned n represents how much of the input we where able to write,
// this may be different than the actual number of bytes written since
// Converter.Encode converts multibyte UTF-8 into singlebyte ISO 8859, i.e. 
// if you write []byte("€€€") using charset ISO_8859_15, it will return 9,
// but it actually just wrote 3 bytes to underlying io.Writer.
func (w *LatinWriter) Write(p []byte) (n int, err error) {

	var e2 error
	var latin []byte
	var size, n2 int

	latin, size, err = w.converter.Encode(p)

	switch err.(type) {

	case UnknownRuneError:
		n = 0
		break

	case UnicodeError:

		// Its either ILLEGAL or PARTIAL (PARTIAL is good)

		if err == ILLEGAL {
			return 0, err
		}

		// Must be PARTIAL

		n2, e2 = w.writer.Write(latin)

		if e2 == nil {
			n = size
		} else {
			n = n2
			err = e2 // priority to inner os.Error
		}

		break

	case nil:
		n, err = w.writer.Write(latin)
		break
	}

	return
}
