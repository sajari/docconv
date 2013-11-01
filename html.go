package main

import (
	"io"
	"log"
	"bytes"
)

// Convert HTML
func ConvertHtml(input io.Reader) (string, map[string]string) {
	meta := make(map[string]string)
	cleanXml, err := Tidy(input, false)
	if err != nil {
		log.Println("Tidy:", err)
	}
	return Html2Text(bytes.NewReader(cleanXml)), meta
}

func Html2Text(input io.Reader) string {
	return Xml2Text(input, []string{ "br", "p", "h1", "h2", "h3", "h4" }, []string{}, false)
}
