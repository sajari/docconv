// +build !ocr

package docconv

import (
	"fmt"
	"io"
)

func ConvertImage(r io.Reader) (string, map[string]string, error) {
	return "", nil, fmt.Errorf("docconv not built with `ocr` build tag")
}

func SetImageLanguages(string) {}
