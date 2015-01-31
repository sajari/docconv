package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
)

// Convert XML input
func ConvertXml(input io.Reader) (string, map[string]string) {
	meta := make(map[string]string)
	cleanXml, err := Tidy(input, true)
	if err != nil {
		log.Println("Tidy:", err)
	}
	return Xml2Text(bytes.NewReader(cleanXml), []string{}, []string{}, true), meta
}

// Convert XML to plain text given how to treat elements
func Xml2Text(input io.Reader, breaks []string, skip []string, strict bool) string {
	var output string
	var depth int

	x := xml.NewDecoder(input)
	x.Strict = strict
	for d, err := x.Token(); d != nil && err == nil; d, err = x.Token() {
		switch v := d.(type) {
		case xml.CharData:
			output += string(v)
		case xml.StartElement:
			for _, breakElement := range breaks {
				if v.Name.Local == breakElement {
					output += "\n"
				}
			}
			for _, skipElement := range skip {
				if v.Name.Local == skipElement {
					depth = 1
					for d, err := x.Token(); d != nil && err == nil; d, err = x.Token() {
						switch d.(type) {
						case xml.StartElement:
							depth++
						case xml.EndElement:
							depth--
						}
						if depth == 0 {
							break
						}
					}
				}
			}
		}
	}
	return output
}

// Convert XML to a nested string map
func Xml2Map(input io.Reader) map[string]string {
	output := make(map[string]string)
	x := xml.NewDecoder(input)
	var tagName string
	for d, err := x.Token(); d != nil && err == nil; d, err = x.Token() {
		switch v := d.(type) {
		case xml.StartElement:
			tagName = string(v.Name.Local)
		case xml.CharData:
			output[tagName] = string(v)
		}
	}
	return output
}
