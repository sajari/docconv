package goose

import (
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/fatih/set.v0"
)

type Article struct {
	Title           string
	CleanedText     string
	MetaDescription string
	MetaLang        string
	MetaFavicon     string
	MetaKeywords    string
	CanonicalLink   string
	Domain          string
	TopNode         *goquery.Selection
	TopImage        string
	Tags            *set.Set
	Movies          *set.Set
	FinalUrl        string
	LinkHash        string
	RawHtml         string
	Doc             *goquery.Document
	//raw_doc
	PublishDate    string
	AdditionalData map[string]string
	Delta          int64
}

//Simple ToString: it shows just the title
//TODO: add more fields and pretty print
func (article *Article) ToString() string {
	out := article.Title
	return out
}
