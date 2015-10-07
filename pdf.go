package docconv

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func compareExt(ext string, exts []string) bool {
	for _, e := range exts {
		if ext == e {
			return true
		}
	}
	return false
}

func PdfImages(path string) (string, map[string]string, error) {
	tmp, err := ioutil.TempDir("/tmp", "tmp-imgs-")
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	tmpDir := fmt.Sprintf("%s/", tmp)
	defer os.RemoveAll(tmpDir)

	_, err = exec.Command("pdfimages", "-j", path, tmpDir).Output()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	files := []string{}
	m := make(map[int]string)

	walkFunc := func(path string, info os.FileInfo, err error) error {
		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}

		exts := []string{".jpg", ".tif", ".tiff", ".png"}
		if compareExt(filepath.Ext(path), exts) == true {
			files = append(files, path)
		}
		return nil
	}
	filepath.Walk(tmpDir, walkFunc)

	var wg sync.WaitGroup
	wg.Add(len(files))
	for indx, p := range files {
		go func(idx int, pathFile string, m map[int]string, ww *sync.WaitGroup) {
			defer ww.Done()
			f, err := os.Open(pathFile)
			if err != nil {
				log.Println(err)
			}
			out, _, _ := ConvertImage(f)
			m[idx] = out
			f.Close()
		}(indx, p, m, &wg)
	}
	wg.Wait()

	o := make([]string, len(m))
	for i := 0; i < len(m); i++ {
		o = append(o, m[i])
	}
	return strings.Join(o, " "), nil, nil
}

// PdfHasImage verify if `path` (PDF) has images
func PdfHasImage(path string) bool {
	cmd := "pdffonts -l 5 %s | tail -n +3 | cut -d' ' -f1 | sort | uniq"
	out, err := exec.Command("bash", "-c", fmt.Sprintf(cmd, path)).Output()
	if err != nil {
		log.Println(err)
		return false
	}
	if string(out) == "" {
		return true
	}
	return false
}

// Convert PDF
func ConvertPDF(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Verify if pdf has images or is pdf only-text
	if PdfHasImage(f.Name()) == true {
		return PdfImages(f.Name())
	}

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
