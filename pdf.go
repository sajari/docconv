package docconv

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// ConvertPDF converts a PDF file to text.
func ConvertPDF(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Meta data
	mc := make(chan map[string]string)
	go func() {
		meta := make(map[string]string)
		metaStr, err := exec.Command("pdfinfo", f.Name()).Output()
		if err != nil {
			// TODO: Remove this.
			log.Println("pdfinfo:", err)
		}

		// Parse meta output
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, ":", 2); len(parts) > 1 {
				meta[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed dates into unix timestamps
		if tmp, ok := meta["ModDate"]; ok {
			if t, err := time.Parse(time.ANSIC, tmp); err == nil {
				meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if tmp, ok := meta["CreationDate"]; ok {
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
		if len(body) == 0 {

			// grab the metadata, mitm-style
			meta := <-mc
			go func() { mc <- meta }()

			pages, err := strconv.Atoi(meta["Pages"])
			if err != nil {
				log.Printf("failed to get number of pages from '%s': %v", meta["pages"], err)
				return
			}

			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Fatal(err)
			}

			for page := 0; page < pages; page++ {
				tmpImageFileName := fmt.Sprintf("%s-%d.jpg", f.Name(), page)
				args := []string{
					fmt.Sprintf("-dFirstPage=%d", page+1),
					fmt.Sprintf("-dLastPage=%d", page+1),
					"-dNOGC",
					// "-dQUIET",
					// "-dBandHeight=100",
					// "-dBandBufferSpace=500000000",
					// "-dBufferSpace=1000000000",
					// "-sBandListStorage=memory",
					// "-dNumRenderingThreads=2",
					"-sDEVICE=jpeg",
					"-r300",
					"-o", tmpImageFileName,
					"-f", filepath.Join(dir, f.Name()),
				}

				_, err := exec.Command("gs", args...).Output()
				if err != nil {
					log.Printf("failed to execute gs: %v", err)
					return
				}

				// stdout, err := cmd.StdoutPipe()
				// if err != nil {
				// 	log.Printf("failed to get stdout from gs call: %v", err)
				// 	return
				// }

				tmpImageFile, err := os.Open(tmpImageFileName)
				if err != nil {
					log.Printf("failed to open temp image file: %v", err)
					return
				}
				defer func() {
					err := tmpImageFile.Close()
					if err != nil {
						log.Printf("failed to close tmpImageFile: %v", err)
					}
					err = os.Remove(tmpImageFileName)
					if err != nil {
						log.Printf("failed to remove tmpImageFileName: %v", err)
					}
				}()

				result, _, err := ConvertImage(tmpImageFile)
				if err != nil {
					log.Printf("failed to run ConvertImage on gs output: %v", err)
					return
				}
				bc <- result
			}
		}
		bc <- string(body)
	}()

	return <-bc, <-mc, nil
}
