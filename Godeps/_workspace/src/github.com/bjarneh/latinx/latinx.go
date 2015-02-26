// Copyright 2011 bjarneh@ifi.uio.no. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package latinx

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// Constants used to fetch *Converter
const (
	ISO_8859_1 = iota
	ISO_8859_2
	ISO_8859_3
	ISO_8859_4
	ISO_8859_5
	ISO_8859_6
	ISO_8859_7
	ISO_8859_8
	ISO_8859_9
	ISO_8859_10
	ISO_8859_11
	ISO_8859_13
	ISO_8859_14
	ISO_8859_15
	ISO_8859_16
	// Extended Latin-1 (Windows only, not a standard)
	Windows1252
	// Common aliases for the standards
	Latin1   = ISO_8859_1
	Latin2   = ISO_8859_2
	Latin3   = ISO_8859_3
	Latin4   = ISO_8859_4
	Cyrillic = ISO_8859_5
	Arabic   = ISO_8859_6
	Greek    = ISO_8859_7
	Hebrew   = ISO_8859_8
	Latin5   = ISO_8859_9
	Latin6   = ISO_8859_10
	Thai     = ISO_8859_11
	Latin7   = ISO_8859_13
	Latin8   = ISO_8859_14
	Latin9   = ISO_8859_15
	Latin10  = ISO_8859_16
	// The numbers (1,2) are just meant to 
	// distinguish PARTIAL from ILLEGAL
	PARTIAL = UnicodeError(1)
	ILLEGAL = UnicodeError(2)
)

var converters map[int]*Converter

// A Converter holds mappings from ISO 8859 => UTF-8, and vice verca.
type Converter struct {
	id          string
	latinToUtf8 map[byte][]byte
	utf8ToLatin map[int]byte
}

func (c *Converter) String() string {
	return c.id
}

// Error type for Partial UTF-8 sequences
type UnicodeError int

func (e UnicodeError) Error() string {
	if int(e) == 1 {
		return "Partial UTF-8 sequence"
	}
	return "Illegal UTF-8 sequence"
}

// Error type for unknown ISO 8859 byte
type UnknownByteError string

func (e UnknownByteError) Error() string {
	return string(e)
}

// Error type for unknown UTF-8 runes
type UnknownRuneError string

func (e UnknownRuneError) Error() string {
	return string(e)
}

func init() {

	converters = make(map[int]*Converter)

	converters[ISO_8859_1] = newISO_8859_1Converter()
	converters[ISO_8859_2] = newISO_8859_2Converter()
	converters[ISO_8859_3] = newISO_8859_3Converter()
	converters[ISO_8859_4] = newISO_8859_4Converter()
	converters[ISO_8859_5] = newISO_8859_5Converter()
	converters[ISO_8859_6] = newISO_8859_6Converter()
	converters[ISO_8859_7] = newISO_8859_7Converter()
	converters[ISO_8859_8] = newISO_8859_8Converter()
	converters[ISO_8859_9] = newISO_8859_9Converter()
	converters[ISO_8859_10] = newISO_8859_10Converter()
	converters[ISO_8859_11] = newISO_8859_11Converter()
	converters[ISO_8859_13] = newISO_8859_13Converter()
	converters[ISO_8859_14] = newISO_8859_14Converter()
	converters[ISO_8859_15] = newISO_8859_15Converter()
	converters[ISO_8859_16] = newISO_8859_16Converter()
	converters[Windows1252] = newWindows1252Converter()

}

// Return *Converter || nil for unknown charset
func Get(charset int) *Converter {
	conv, ok := converters[charset]
	if !ok {
		return nil
	}
	return conv
}

// Return the String representation of all available encodings
func Available() (all []string) {
	for _, v := range converters {
		all = append(all, v.String())
	}
	return
}

// Convert a UTF-8 byte sequence into a ISO 8859 byte sequence. The errors returned
// by this function are either UnicodeError, which means that a partial UTF-8 symbol
// or an illegal UTF-8 sequence was found, i.e. either latinx.ILLEGAL, or latinx.PARTIAL.
// When a UnicodeError is returned, success < len(utf_8), and success indicates how
// many bytes that was successfully converted into UTF-8 bytes.
// If this function returns an UnknownRuneError, it means that the charset of the
// Converter has no mapping for a rune (UTF-8 letter) found in the utf_8 array.
func (c *Converter) Encode(utf_8 []byte) (latin []byte, success int, err error) {

	var ok bool
	var latinByte byte
	var offset, size int
	var rne rune
	var errmsg string
	var buf *bytes.Buffer

	buf = bytes.NewBuffer(make([]byte, len(utf_8)))
	buf.Reset()

	for offset < len(utf_8) {

		rne, size = utf8.DecodeRune(utf_8[offset:])

		if rne == utf8.RuneError {
			if utf8.RuneStart(utf_8[offset]) && len(utf_8)-offset < utf8.UTFMax {
				return buf.Bytes(), offset, PARTIAL // UnicodeError
			} else {
				return buf.Bytes(), offset, ILLEGAL // UnicodeError
			}
		} else if rne < utf8.RuneSelf {
			buf.WriteByte(utf_8[offset])
			offset++
		} else {
			latinByte, ok = c.utf8ToLatin[int(rne)]
			if !ok {
				errmsg = fmt.Sprintf("undefined: 0x%X in %s", rne, c.id)
				err = UnknownRuneError(errmsg)
				return buf.Bytes(), offset, err
			}
			buf.WriteByte(latinByte)
			offset += size
		}
	}
	return buf.Bytes(), offset, err
}

// Convert a ISO 8859 byte sequence into a UTF-8 byte sequence.
// If this function returns a UnknownByteError, the charset of the
// Converter does not have a unicode mapping for a byte found in latin.
func (c *Converter) Decode(latin []byte) (utf_8 []byte, err error) {

	var offset, i int
	var ok bool
	var utf8symbol []byte
	var errmsg string
	var buf *bytes.Buffer

	buf = bytes.NewBuffer(make([]byte, len(latin)*2))
	buf.Reset()

	for offset < len(latin) {

		if latin[offset] < utf8.RuneSelf {
			buf.WriteByte(latin[offset])
			offset++
		} else {
			utf8symbol, ok = c.latinToUtf8[latin[offset]]
			if !ok {
				errmsg = fmt.Sprintf("0x%X not valid in: %s", latin[offset], c.id)
				err = UnknownByteError(errmsg)
				return buf.Bytes(), err
			}
			for i = 0; i < len(utf8symbol); i++ {
				buf.WriteByte(utf8symbol[i])
			}
			offset++
		}
	}
	return buf.Bytes(), err
}

// Wrappers to avoid fetching a *Converter to do a single encode/decode

// Convert a UTF-8 encoded slice to a ISO-8859 encoded slice
func Encode(charset int, utf_8 []byte) (latin []byte, success int, err error) {
	return converters[charset].Encode(utf_8)
}

// Convert a ISO-8859 encoded slice to a UTF-8 encoded slice
func Decode(charset int, latin []byte) (utf_8 []byte, err error) {
	return converters[charset].Decode(latin)
}
