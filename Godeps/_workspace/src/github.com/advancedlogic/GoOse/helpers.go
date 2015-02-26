package goose

import (
	"crypto/md5"
	"fmt"
	"github.com/bjarneh/latinx"
	"io"
	"strings"
	"unicode/utf8"
)

type Helper struct {
	urlString string
	url       string
	linkHash  string
}

func NewRawHelper(url string, rawHtml string) Helper {
	if !utf8.ValidString(rawHtml) {
		converter := latinx.Get(latinx.ISO_8859_1)
		rawHtmlBytes, err := converter.Decode([]byte(rawHtml))
		if err != nil {
			fmt.Println(err.Error())
		}
		rawHtml = string(rawHtmlBytes)
	}

	h := md5.New()
	io.WriteString(h, url)
	bytes := h.Sum(nil)
	helper := Helper{
		urlString: url,
		url:       url,
		linkHash:  fmt.Sprintf("%s.%d", string(bytes), TimeInNanoseconds()),
	}
	return helper
}

func NewUrlHelper(url string) Helper {
	finalUrl := ""
	if strings.Contains(url, "#!") {
		finalUrl = strings.Replace(url, "#!", "?_escaped_fragment_=", -1)
	} else {
		finalUrl = url
	}
	h := md5.New()
	io.WriteString(h, finalUrl)
	bytes := h.Sum(nil)
	helper := Helper{
		urlString: finalUrl,
		url:       finalUrl,
		linkHash:  fmt.Sprintf("%s.%d", string(bytes), TimeInNanoseconds()),
	}
	return helper
}
