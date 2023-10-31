//go:build appengine

package docconv

import (
	"io"
)

func HTMLReadability(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
