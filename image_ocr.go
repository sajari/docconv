// +build ocr

package docconv

import (
	"fmt"
	"io"
	"sync"

	"github.com/otiai10/gosseract"
)

var languages = struct {
	sync.RWMutex
	lang string
}{lang: "eng"}

func ConvertImage(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	meta := make(map[string]string)
	out := make(chan string, 1)
	go func(file *LocalFile) {
		languages.RLock()
		body := gosseract.Must(gosseract.Params{Src: file.Name(), Languages: languages.lang})
		languages.RUnlock()
		out <- string(body)
	}(f)

	return <-out, meta, nil
}

func SetLanguages(l string) {
	languages.Lock()
	languages.lang = l
	languages.Unlock()
}
