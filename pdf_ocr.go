// +build ocr

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

func cleanupTemp(tmpDir string) {
	err := os.RemoveAll(tmpDir)
	if err != nil {
		log.Println(err)
	}
}

func ConvertPDFImages(path string) (BodyResult, error) {
	bodyResult := BodyResult{}

	tmp, err := ioutil.TempDir("/tmp", "tmp-imgs-")
	if err != nil {
		bodyResult.err = err
		return bodyResult, err
	}
	tmpDir := fmt.Sprintf("%s/", tmp)

	defer cleanupTemp(tmpDir)

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

	if len(files) < 1 {
		return bodyResult, nil
	}

	var wg sync.WaitGroup

	data := make(chan string, 1)

	wg.Add(len(files))
	for _, p := range files {
		go func(pathFile string) {

			f, err := os.Open(pathFile)

			if err != nil {
				bodyResult.err = err
			}

			defer f.Close()
			out, _, err := ConvertImage(f)
			if err != nil {
				bodyResult.err = err
			}

			wg.Done()

			data <- out

		}(p)
	}

	wg.Wait()

	go func() {
		close(data)
	}()

	for str := range data {
		bodyResult.body += str + " "
	}

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

func ConvertPDF(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	bodyResult, metaResult, textConvertErr := ConvertPDFText(f.Name())
	if textConvertErr != nil {
		return "", nil, textConvertErr
	}
	if bodyResult.err != nil {
		return "", nil, bodyResult.err
	}
	if metaResult.err != nil {
		return "", nil, metaResult.err
	}

	if !PDFHasImage(f.Name()) {
		return bodyResult.body, metaResult.meta, nil
	}

	fmt.Println("converting PDF images")

	imageConvertResult, imageConvertErr := ConvertPDFImages(f.Name())
	if imageConvertErr != nil {
		log.Println(imageConvertErr)
		return bodyResult.body, metaResult.meta, nil
	}
	if imageConvertResult.err != nil {
		log.Println(imageConvertResult.err)
		return bodyResult.body, metaResult.meta, nil
	}

	fmt.Println("no errors...")

	fullBody := strings.Join([]string{bodyResult.body, imageConvertResult.body}, " ")

	fmt.Println("returning response...")

	return fullBody, metaResult.meta, nil

}
