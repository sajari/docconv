// +build ocr

package docconv

import (
	"fmt"
	"io"
	"log"
)

func ConvertPDF(r io.Reader) (string, map[string]string, error) {
	log.Println("checking if pdf has image.... ")
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Verify if pdf has images or is pdf only-text
	log.Println("checking if pdf has image... ")
	if PDFHasImage(f.Name()) {
		log.Println("pdf does have image")
		bodyResult, imageConvertErr := ConvertImagePDF(f.Name())
		if imageConvertErr != nil {
			return "", nil, imageConvertErr
		}
		if bodyResult.err != nil {
			return "", nil, bodyResult.err
		}

		return bodyResult.body, nil, nil
	}
	bodyResult, metaResult, textConvertErr := ConvertTextPDF(f.Name())
	if textConvertErr != nil {
		return "", nil, textConvertErr
	}
	if bodyResult.err != nil {
		return "", nil, bodyResult.err
	}
	if metaResult.err != nil {
		return "", nil, metaResult.err
	}
	return bodyResult.body, metaResult.meta, nil

}
