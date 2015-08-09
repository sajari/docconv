package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"
)

// Convert ODT to text
func ConvertODT(r io.Reader) (string, map[string]string) {
	meta := make(map[string]string)
	var textBody string

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("ioutil.ReadAll:", err)
		return "", nil
	}
	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		log.Println("zip.NewReader:", err)
		return "", nil
	}

	for _, f := range zr.File {
		if f.Name == "meta.xml" {
			rc, _ := f.Open()
			defer rc.Close()
			info := XMLToMap(rc)
			if tmp, ok := info["creator"]; ok {
				meta["Author"] = tmp
			}
			if tmp, ok := info["date"]; ok {
				if t, err := time.Parse("2006-01-02T15:04:05", tmp); err == nil {
					meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
				}
			}
			if tmp, ok := info["creation-date"]; ok {
				if t, err := time.Parse("2006-01-02T15:04:05", tmp); err == nil {
					meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
				}
			}
		} else if f.Name == "content.xml" {
			rc, _ := f.Open()
			defer rc.Close()
			textBody = XMLToText(rc, []string{"br", "p", "tab"}, []string{}, true)
		}
	}

	return textBody, meta
}
