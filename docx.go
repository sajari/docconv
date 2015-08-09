package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"time"
)

// Convert DOCX to text
func ConvertDocx(r io.Reader) (string, map[string]string) {
	meta := make(map[string]string)
	var textHeader, textBody, textFooter string

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

	// Regular expression for XML files to include in the text parsing
	reHeaderFile, _ := regexp.Compile("^word/header[0-9]+.xml$")
	reFooterFile, _ := regexp.Compile("^word/footer[0-9]+.xml$")

	for _, f := range zr.File {
		if f.Name == "docProps/core.xml" {
			rc, _ := f.Open()
			defer rc.Close()
			info := XMLToMap(rc)
			if tmp, ok := info["modified"]; ok {
				if t, err := time.Parse(time.RFC3339, tmp); err == nil {
					meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
				}
			}
			if tmp, ok := info["created"]; ok {
				if t, err := time.Parse(time.RFC3339, tmp); err == nil {
					meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
				}
			}
		} else if f.Name == "word/document.xml" {
			rc, _ := f.Open()
			defer rc.Close()
			textBody = DocxXMLToText(rc)
		} else if reHeaderFile.MatchString(f.Name) {
			rc, _ := f.Open()
			defer rc.Close()
			textHeader += DocxXMLToText(rc) + "\n"
		} else if reFooterFile.MatchString(f.Name) {
			rc, _ := f.Open()
			defer rc.Close()
			textFooter += DocxXMLToText(rc) + "\n"
		}
	}

	return textHeader + "\n" + textBody + "\n" + textFooter, meta
}

func DocxXMLToText(r io.Reader) string {
	return XMLToText(r, []string{"br", "p", "tab"}, []string{"instrText", "script"}, true)
}
