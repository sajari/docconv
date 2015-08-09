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
		// TODO: Return error
		log.Println("Tidy:", err)
	}
	return Xml2Text(bytes.NewReader(cleanXml), []string{}, []string{}, true), meta
}

// Convert XML to plain text given how to treat elements
func Xml2Text(r io.Reader, breaks []string, skip []string, strict bool) string {
	var result string

	dec := xml.NewDecoder(r)
	dec.Strict = strict
	for {
		t, err := dec.Token()
		if err != nil {
			// TODO: Handle non io.EOF error
			break
		}

		switch v := t.(type) {
		case xml.CharData:
			result += string(v)
		case xml.StartElement:
			for _, breakElement := range breaks {
				if v.Name.Local == breakElement {
					result += "\n"
				}
			}
			for _, skipElement := range skip {
				if v.Name.Local == skipElement {
					depth := 1
					for {
						t, err := dec.Token()
						if err != nil {
							// TODO: Handle non io.EOF error
							break
						}

						switch t.(type) {
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
	return result
}

// Convert XML to a nested string map
func Xml2Map(r io.Reader) map[string]string {
	m := make(map[string]string)
	dec := xml.NewDecoder(r)
	var tagName string
	for {
		t, err := dec.Token()
		if err != nil {
			break
		}

		switch v := t.(type) {
		case xml.StartElement:
			tagName = string(v.Name.Local)
		case xml.CharData:
			m[tagName] = string(v)
		}
	}
	return m
}
