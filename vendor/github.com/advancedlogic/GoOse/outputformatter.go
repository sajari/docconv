package goose

import (
	"github.com/advancedlogic/goquery"
	"golang.org/x/net/html"
	"regexp"
	"strconv"
	"strings"
)

var normalizeWhitespaceRegexp = regexp.MustCompile(`[ \r\f\v\t]+`)
var normalizeNl = regexp.MustCompile(`[\n]+`)

type outputFormatter struct {
	topNode  *goquery.Selection
	config   Configuration
	language string
}

func (formatter *outputFormatter) getLanguage(article *Article) string {
	if formatter.config.useMetaLanguage {
		if article.MetaLang != "" {
			return article.MetaLang
		}
	}
	return formatter.config.targetLanguage
}

func (formatter *outputFormatter) getTopNode() *goquery.Selection {
	return formatter.topNode
}

func (formatter *outputFormatter) getFormattedText(article *Article) (output string, links []string) {
	formatter.topNode = article.TopNode
	formatter.language = formatter.getLanguage(article)
	if formatter.language == "" {
		formatter.language = formatter.config.targetLanguage
	}
	formatter.removeNegativescoresNodes()
	links = formatter.linksToText()
	formatter.replaceTagsWithText()
	formatter.removeParagraphsWithFewWords()
	output = formatter.getOutputText()
	return
}

func (formatter *outputFormatter) convertToText() string {
	var txts []string
	selections := formatter.topNode
	selections.Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		if txt != "" {
			// txt = txt //unescape
			txtLis := strings.Trim(txt, "\n")
			txts = append(txts, txtLis)
		}
	})
	return strings.Join(txts, "\n\n")
}

func (formatter *outputFormatter) linksToText() []string {
	urlList := []string{}
	links := formatter.topNode.Find("a")
	links.Each(func(i int, a *goquery.Selection) {
		imgs := a.Find("img")
		if imgs.Length() == 0 {
			node := a.Get(0)
			node.Data = a.Text()
			node.Type = html.TextNode
			// save a list of URLs
			url, _ := a.Attr("href")
			isValidUrl, _ := regexp.MatchString("^http[s]?://", url)
			if isValidUrl {
				urlList = append(urlList, url)
			}
		}
	})
	return urlList
}

func (formatter *outputFormatter) getOutputText() string {

	out := formatter.topNode.Text()
	out = normalizeWhitespaceRegexp.ReplaceAllString(out, " ")

	strArr := strings.Split(out, "\n")
	resArr := []string{}

	for i, v := range strArr {
		v = strings.TrimSpace(v)
		if v != "" {
			resArr = append(resArr, v)
		} else if i > 2 && strArr[i-2] != "" {
			resArr = append(resArr, "")
		}
	}

	out = strings.Join(resArr, "\n")
	out = normalizeNl.ReplaceAllString(out, "\n\n")

	out = strings.TrimSpace(out)
	return out
}

func (formatter *outputFormatter) removeNegativescoresNodes() {
	gravityItems := formatter.topNode.Find("*[gravityScore]")
	gravityItems.Each(func(i int, s *goquery.Selection) {
		score := 0
		sscore, exists := s.Attr("gravityScore")
		if exists {
			score, _ = strconv.Atoi(sscore)
			if score < 1 {
				sNode := s.Get(0)
				sNode.Parent.RemoveChild(sNode)
			}
		}

	})
}

func (formatter *outputFormatter) replaceTagsWithText() {
	strongs := formatter.topNode.Find("strong")
	strongs.Each(func(i int, strong *goquery.Selection) {
		text := strong.Text()
		node := strong.Get(0)
		node.Type = html.TextNode
		node.Data = text
	})

	bolds := formatter.topNode.Find("b")
	bolds.Each(func(i int, bold *goquery.Selection) {
		text := bold.Text()
		node := bold.Get(0)
		node.Type = html.TextNode
		node.Data = text
	})

	italics := formatter.topNode.Find("i")
	italics.Each(func(i int, italic *goquery.Selection) {
		text := italic.Text()
		node := italic.Get(0)
		node.Type = html.TextNode
		node.Data = text
	})
}

func (formatter *outputFormatter) removeParagraphsWithFewWords() {
	language := formatter.language
	if language == "" {
		language = "en"
	}
	allNodes := formatter.topNode.Children()
	allNodes.Each(func(i int, s *goquery.Selection) {
		sw := formatter.config.stopWords.stopWordsCount(language, s.Text())
		if sw.wordCount < 5 && s.Find("object").Length() == 0 && s.Find("em").Length() == 0 {
			node := s.Get(0)
			node.Parent.RemoveChild(node)
		}
	})

}
