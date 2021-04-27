// +build appengine

package docconv

import (
	"io"
	"io/ioutil"
)

func HTMLReadability(r io.Reader) []byte {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		logger.Printf("HTMLReadability: %v", err)
		return nil
	}
	return b
}
