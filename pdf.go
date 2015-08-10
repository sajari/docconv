package docconv

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"
)

// Convert PDF
func ConvertPDF(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Meta data
	mc := make(chan map[string]string, 1)
	go func() {
		meta := make(map[string]string)
		metaStr, err := exec.Command("pdfinfo", f.Name()).Output()
		if err != nil {
			// TODO: Remove this.
			log.Println("pdfinfo:", err)
		}

		// Parse meta output
		info := make(map[string]string)
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, ":", 2); len(parts) > 1 {
				info[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed meta
		if tmp, ok := info["Author"]; ok {
			meta["Author"] = tmp
		}
		if tmp, ok := info["ModDate"]; ok {
			if t, err := time.Parse(time.ANSIC, tmp); err == nil {
				meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if tmp, ok := info["CreationDate"]; ok {
			if t, err := time.Parse(time.ANSIC, tmp); err == nil {
				meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}

		mc <- meta
	}()

	// Document body
	bc := make(chan string, 1)
	go func() {
		body, err := exec.Command("pdftotext", "-q", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", f.Name(), "-").Output()
		if err != nil {
			// TODO: Remove this.
			log.Println("pdftotext:", err)
		}
		bc <- string(body)
	}()

	return <-bc, <-mc, nil
}
