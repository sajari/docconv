package docconv

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type meta struct {
	m map[string]string
	e error
}

type body struct {
	b string
	e error
}

// ConvertDoc converts an MS Word .doc to text.
func ConvertDoc(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r)
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Meta data
	mc := make(chan meta, 1)
	go func() {
		metaData := make(map[string]string)
		metaStr, err := exec.Command("wvSummary", f.Name()).Output()
		if err != nil {
			err = errors.WithMessage(err, "warn: wxSummary error")
		}

		// Parse metaData output
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, "=", 2); len(parts) > 1 {
				metaData[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed metaData
		if tmp, ok := metaData["Last Modified"]; ok {
			t, timeParseErr := time.Parse(time.RFC3339, tmp)
			switch timeParseErr {
			case nil:
				metaData["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			default:
				err = errors.WithMessage(err, "warn: time.Parse of \"Last Modified\" metaData key value failed")
			}
		}
		if tmp, ok := metaData["Created"]; ok {
			t, timeParseErr := time.Parse(time.RFC3339, tmp)
			switch timeParseErr {
			case nil:
				metaData["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			default:
				err = errors.WithMessage(err, "warn: time.Parse of \"Created\" metaData key value failed")
			}
		}

		mc <- meta{
			m: metaData,
			e: err,
		}
	}()

	// Document body
	bc := make(chan body, 1)
	go func() {

		// Save output to a file
		outputFile, err := ioutil.TempFile("/tmp", "sajari-convert-")
		if err != nil {
			bc <- body{
				e: errors.WithMessage(err, "fatal: failed to create tempfile"),
			}
			return
		}
		defer os.Remove(outputFile.Name())

		wvTextErr := exec.Command("wvText", f.Name(), outputFile.Name()).Run()
		if wvTextErr != nil {
			bc <- body{
				e: errors.WithMessage(wvTextErr, "fatal: wvText error"),
			}
			return
		}

		var buf bytes.Buffer
		_, bufErr := buf.ReadFrom(outputFile)
		if bufErr != nil {
			bc <- body{
				e: errors.WithMessage(bufErr, "fatal: wvText error reading from tempfile"),
			}
			return
		}

		bc <- body{
			b: buf.String(),
		}
	}()

	// TODO: Should errors in either of the above Goroutines stop things from progressing?
	// The errors need to be aggregated into Warn or Fatal groups, fail on Fatal errors
	// but, bubble up warnings to inform the user without blocking content
	body := <-bc
	meta := <-mc

	// TODO: Check for errors instead of len(body) == 0?
	// Since this is catching an empty content, the document may actually be empty but there could have been an error.
	// Having an empty document is valid and we are returning errors in each case they can occur, we check for errors.
	switch errors.Cause(body.e) {
	case nil:
		// carry on
	case os.ErrClosed, os.ErrPermission:
		// encountered an os error while attempting to create or read tempfile
		return body.b, meta.m, body.e
	default:
		switch errors.Cause(body.e).(type) {
		case *os.PathError:
			return body.b, meta.m, body.e
		}
		// Falling back to attempt decoding as "docx", discarding errors received so far...
		// We definitely want to do this in case of &exec.ExitError{}, but then what should be the default behavior?
		// this should be fine for now...
		f.Seek(0, 0)
		return ConvertDocx(f)
	}
	return body.b, meta.m, body.e
}
