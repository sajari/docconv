package docconv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	exts = []string{".jpg", ".tif", ".tiff", ".png", ".pbm"}
)

func compareExt(ext string, exts []string) bool {
	for _, e := range exts {
		if ext == e {
			return true
		}
	}
	return false
}

// Meta data
type MetaResult struct {
	meta map[string]string
	err  error
}

type BodyResult struct {
	body string
	err  error
}

func ConvertImagePDF(path string) (BodyResult, error) {
	bodyResult := BodyResult{}

	tmp, err := ioutil.TempDir("/tmp", "tmp-imgs-")
	if err != nil {
		log.Println(err)
		return bodyResult, err
	}
	tmpDir := fmt.Sprintf("%s/", tmp)
	defer os.RemoveAll(tmpDir)

	_, err = exec.Command("pdfimages", "-j", path, tmpDir).Output()
	if err != nil {
		return bodyResult, err
	}

	files := []string{}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}

		if compareExt(filepath.Ext(path), exts) {
			files = append(files, path)
		}
		return nil
	}
	filepath.Walk(tmpDir, walkFunc)

	var wg sync.WaitGroup

	m := struct {
		sync.Mutex
		fileMap map[int]string
	}{}

	m.fileMap = make(map[int]string)

	wg.Add(len(files))
	for indx, p := range files {
		go func(idx int, pathFile string, ww *sync.WaitGroup) {
			defer ww.Done()
			f, err := os.Open(pathFile)
			if err != nil {
				log.Println(err)
			}
			out, _, err := ConvertImage(f)
			if err != nil {
				bodyResult.err = err
			}

			m.Lock()
			m.fileMap[idx] = out
			m.Unlock()

			f.Close()
		}(indx, p, &wg)
	}
	wg.Wait()

	o := make([]string, len(m.fileMap))

	for i := 0; i < len(m.fileMap); i++ {
		o = append(o, m.fileMap[i])
	}

	bodyResult.body = strings.Join(o, " ")

	return bodyResult, nil
}

// PdfHasImage verify if `path` (PDF) has images
func PDFHasImage(path string) bool {
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

func ConvertTextPDF(path string) (BodyResult, MetaResult, error) {
	metaResult := MetaResult{}
	bodyResult := BodyResult{}
	mr := make(chan MetaResult, 1)
	go func() {
		metaResult.meta = make(map[string]string)
		metaStr, err := exec.Command("pdfinfo", path).Output()
		if err != nil {
			metaResult.err = err
			mr <- metaResult
			return
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
			metaResult.meta["Author"] = tmp
		}
		if tmp, ok := info["ModDate"]; ok {
			if t, err := time.Parse(time.ANSIC, tmp); err == nil {
				metaResult.meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if tmp, ok := info["CreationDate"]; ok {
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
