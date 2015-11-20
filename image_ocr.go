// +build ocr

package docconv

import (
	"fmt"
	"io"

	"github.com/otiai10/gosseract"
)

var languages = "eng"

func ConvertImage(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	meta := make(map[string]string)
	out := make(chan string, 1)
	go func(file *LocalFile) {
		fmt.Println(languages)
		body := gosseract.Must(gosseract.Params{Src: file.Name(), Languages: languages})
		out <- string(body)
	}(f)

	return <-out, meta, nil
}

func SetLanguages(l string) {
	languages = l
}

