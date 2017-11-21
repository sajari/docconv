package docconv

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Meta data
type MetaResult struct {
	meta map[string]string
	err  error
}

type BodyResult struct {
	body string
	err  error
}

// Convert PDF

func ConvertPDFText(path string) (BodyResult, MetaResult, error) {
	metaResult := MetaResult{meta: make(map[string]string)}
	bodyResult := BodyResult{}
	mr := make(chan MetaResult, 1)
	go func() {
		metaStr, err := exec.Command("pdfinfo", path).Output()
		if err != nil {
			metaResult.err = err
			mr <- metaResult
			return
		}

		// Parse meta output
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, ":", 2); len(parts) > 1 {
				metaResult.meta[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed meta
		if tmp, ok := metaResult.meta["Author"]; ok {
			metaResult.meta["Author"] = tmp
		}
		if tmp, ok := metaResult.meta["ModDate"]; ok {
			if t, err := time.Parse(time.ANSIC, tmp); err == nil {
				metaResult.meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if tmp, ok := metaResult.meta["CreationDate"]; ok {
			if t, err := time.Parse(time.ANSIC, tmp); err == nil {
				metaResult.meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}

		mr <- metaResult
	}()

	br := make(chan BodyResult, 1)
	go func() {
		body, err := exec.Command("pdftotext", "-q", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", path, "-").Output()
		if err != nil {
			bodyResult.err = err
		}

		bodyResult.body = string(body)

		br <- bodyResult
	}()

	return <-br, <-mr, nil
}
