package goose

import (
	"crypto/md5"
	"fmt"
	"github.com/bjarneh/latinx"
	"io"
	"strings"
	"time"
	"unicode/utf8"
)

// Helper is a utility struct to clean up URLs and charsets
type Helper struct {
	urlString string
	url       string
	linkHash  string
}

// NewRawHelper converts the text to UTF8
func NewRawHelper(url string, RawHTML string) Helper {
	if !utf8.ValidString(RawHTML) {
		converter := latinx.Get(latinx.ISO_8859_1)
		RawHTMLBytes, err := converter.Decode([]byte(RawHTML))
		if err != nil {
			fmt.Println(err.Error())
		}
		RawHTML = string(RawHTMLBytes)
	}

	h := md5.New()
	io.WriteString(h, url)
	bytes := h.Sum(nil)
	helper := Helper{
		urlString: url,
		url:       url,
		linkHash:  fmt.Sprintf("%s.%d", string(bytes), time.Now().UnixNano()),
	}
	return helper
}

// NewURLHelper wraps the URL
func NewURLHelper(url string) Helper {
	FinalURL := ""
	if strings.Contains(url, "#!") {
		FinalURL = strings.Replace(url, "#!", "?_escaped_fragment_=", -1)
	} else {
		FinalURL = url
	}
	h := md5.New()
	io.WriteString(h, FinalURL)
	bytes := h.Sum(nil)
	helper := Helper{
		urlString: FinalURL,
		url:       FinalURL,
		linkHash:  fmt.Sprintf("%s.%d", string(bytes), time.Now().UnixNano()),
	}
	return helper
}
