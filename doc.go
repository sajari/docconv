package docconv

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/richardlehane/mscfb"
	"github.com/richardlehane/msoleps"
)

// ConvertDoc converts an MS Word .doc to text.
func ConvertDoc(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r)
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Meta data
	mc := make(chan map[string]string, 1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				// TODO: Propagate error.
			}
		}()

		meta := make(map[string]string)

		doc, err := mscfb.New(f)
		if err != nil {
			// TODO: Propagate error.
			mc <- meta
			return
		}

		props := msoleps.New()
		for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {
			if msoleps.IsMSOLEPS(entry.Initial) {
				if err := props.Reset(doc); err != nil {
					// TODO: Propagate error.
					break
				}

				for _, prop := range props.Property {
					meta[prop.Name] = prop.String()
				}
			}
		}

		const defaultTimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

		// Convert parsed meta
		if tmp, ok := meta["LastSaveTime"]; ok {
			if t, err := time.Parse(defaultTimeFormat, tmp); err == nil {
				meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if tmp, ok := meta["CreateTime"]; ok {
			if t, err := time.Parse(defaultTimeFormat, tmp); err == nil {
				meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}

		mc <- meta
	}()

	// Document body
	bc := make(chan string, 1)
	go func() {
		// Save output to a file
		var buf bytes.Buffer
		outputFile, err := os.CreateTemp("/tmp", "sajari-convert-")
		if err != nil {
			bc <- buf.String()
			return
		}
		defer os.Remove(outputFile.Name())

		err = exec.Command("wvText", f.Name(), outputFile.Name()).Run()
		if err != nil {
			// TODO: Propagate error.
		}

		_, err = buf.ReadFrom(outputFile)
		if err != nil {
			// TODO: Propagate error.
		}

		bc <- buf.String()
	}()

	// TODO: Should errors in either of the above Goroutines stop things from progressing?
	body := <-bc
	meta := <-mc

	// TODO: Check for errors instead of len(body) == 0?
	if len(body) == 0 {
		f.Seek(0, 0)
		return ConvertDocx(f)
	}
	return body, meta, nil
}
