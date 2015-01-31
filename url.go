package main

import (
	"bytes"
	"github.com/advancedlogic/GoOse"
	"io"
)

// Convert URL
func ConvertUrl(input io.Reader, readability bool) (string, map[string]string) {
	meta := make(map[string]string)

	buf := new(bytes.Buffer)
	buf.ReadFrom(input)

	g := goose.New()
	article := g.ExtractFromUrl(buf.String())

	meta["title"] = article.Title
	meta["description"] = article.MetaDescription
	meta["image"] = article.TopImage

	return article.CleanedText, meta
}
