/*

Library to encode/decode ISO-8859 byte streams to/from UTF-8.

This library is complete in terms of the ISO 8859 standard,
i.e. all 15 parts are present.

An io.Writer and an io.Reader can be used as well, in order
to write or read ISO-8859 streams from underlying io.Reader/io.Writer.

The Windows-1252 conversion is also included, which uses some
undefined positions in ISO-8859-1 for common characters.


*/
package latinx
