package docconv

import (
	"bytes"
	"github.com/advancedlogic/GoOse"
	"io"
)

// Convert URL
func ConvertURL(input io.Reader, readability bool) (string, map[string]string, error) {
	meta := make(map[string]string)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(input)
	if err != nil {
		return "", nil, err
	}

	g := goose.New()
	article := g.ExtractFromURL(buf.String())

	meta["title"] = article.Title
	meta["description"] = article.MetaDescription
	meta["image"] = article.TopImage

	return article.CleanedText, meta, nil
}
