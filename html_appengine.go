//go:build appengine
// +build appengine

package docconv

import (
	"io"
	"log"
)

func HTMLReadability(r io.Reader) []byte {
	b, err := io.ReadAll(r)
	if err != nil {
		log.Printf("HTMLReadability: %v", err)
		return nil
	}
	return b
}
