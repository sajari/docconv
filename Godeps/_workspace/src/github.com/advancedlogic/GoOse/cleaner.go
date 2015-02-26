package goose

import (
	"container/list"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"log"
	"regexp"
	"strings"
)

type cleaner struct {
	config configuration
}

func NewCleaner(config configuration) cleaner {
	return cleaner{
		config: config,
	}
}

var divToPElementsPattern = regexp.MustCompile("<(a|blockquote|dl|div|img|ol|p|pre|table|ul)")
var tabsRegEx, _ = regexp.Compile("\\t|^\\s+$]")
var REMOVENODES_RE = regexp.MustCompile("" +
	"PopularQuestions|" +
	"[Cc]omentario|" +
	"[Ff]ooter|" +
	"^fn$|" +
	"^inset$|" +
	"^print$|" +
	"^scroll$|" +
	"^side$|" +
	"^side_|" +
	"^widget$|" +
	"ajoutVideo|" +
	"articleheadings|" +
	"author-dropdown|" +
	"blog-pager|" +
	"breadcrumbs|" +
	"byline|" +
	"cabecalho|" +
	"cnnStryHghLght|" +
	"cnn_html_slideshow|" +
	"cnn_strycaptiontxt|" +
	"cnn_strylftcntnt|" +
	"cnn_stryspcvbx|" +
	"combx|" +
	"comment|" +
	"communitypromo|" +
	"contact|" +
	"contentTools2|" +
	"controls|" +
	"^date$|" +
	"detail_new_|" +
	"detail_related_|" +
	"figcaption|" +
	"footnote|" +
	"foot|" +
	"header|" +
	"img_popup_single|" +
	"js_replies|" +
	"[Kk]ona[Ff]ilter|" +
	"leading|" +
	"legende|" +
	"links|" +
	"mediaarticlerelated|" +
	"menucontainer|" +
	"meta$|" +
	"navbar|" +
	"pagetools|" +
	"popup|" +
	"post-attributes|" +
	"post-title|" +
	"relacionado|" +
	"retweet|" +
	"runaroundLeft|" +
	"shoutbox|" +
	"site_nav|" +
	"socialNetworking|" +
	"social_|" +
	"socialnetworking|" +
	"socialtools|" +
	"sponsor|" +
	"sub_nav|" +
	"subscribe|" +
	"tag_|" +
	"tags|" +
	"the_answers|" +
	"timestamp|" +
	"tools|" +
	"vcard|" +
	"welcome_form|" +
	"wp-caption-text")
var CAPTIONS_RE = regexp.MustCompile("^caption$")
var GOOGLE_RE = regexp.MustCompile(" google ")
var MORE_RE = regexp.MustCompile("^[^entry-]more.*$")
var FACEBOOK_RE = regexp.MustCompile("[^-]facebook")
var FACEBOOK_BROADCASTING_RE = regexp.MustCompile("facebook-broadcasting")
var TWITTER_RE = regexp.MustCompile("[^-]twitter")

func (this *cleaner) clean(article *Article) *goquery.Document {
	if this.config.debug {
		log.Println("Starting cleaning phase with Cleaner")
	}
	docToClean := article.Doc
	docToClean = this.cleanArticleTags(docToClean)
	docToClean = this.cleanEMTags(docToClean)
	docToClean = this.dropCaps(docToClean)
	docToClean = this.removeScriptsStyle(docToClean)
	docToClean = this.cleanBadTags(docToClean)
	docToClean = this.cleanFooter(docToClean)
	docToClean = this.cleanAside(docToClean)
	docToClean = this.removeNodesRegEx(docToClean, CAPTIONS_RE)
	docToClean = this.removeNodesRegEx(docToClean, GOOGLE_RE)
	docToClean = this.removeNodesRegEx(docToClean, MORE_RE)
	docToClean = this.removeNodesRegEx(docToClean, FACEBOOK_RE)
	docToClean = this.removeNodesRegEx(docToClean, FACEBOOK_BROADCASTING_RE)
	docToClean = this.removeNodesRegEx(docToClean, TWITTER_RE)
	docToClean = this.cleanParaSpans(docToClean)

	docToClean = this.convertDivsToParagraphs(docToClean, "div")
	docToClean = this.convertDivsToParagraphs(docToClean, "span")
	docToClean = this.convertDivsToParagraphs(docToClean, "article")
	docToClean = this.convertDivsToParagraphs(docToClean, "pre")

	return docToClean
}

func (this *cleaner) cleanArticleTags(doc *goquery.Document) *goquery.Document {
	tags := [3]string{"id", "name", "class"}
	articles := doc.Find("article")
	articles.Each(func(i int, s *goquery.Selection) {
		for _, tag := range tags {
			this.config.parser.delAttr(s, tag)
		}
	})
	return doc
}

func (this *cleaner) cleanEMTags(doc *goquery.Document) *goquery.Document {
	ems := doc.Find("em")
	ems.Each(func(i int, s *goquery.Selection) {
		images := s.Find("img")
		if images.Length() == 0 {
			this.config.parser.dropTag(s)
		}
	})
	if this.config.debug {
		log.Printf("Cleaning %d EM tags\n", ems.Size())
	}
	return doc
}

func (this *cleaner) cleanFooter(doc *goquery.Document) *goquery.Document {
	footer := doc.Find("footer")
	footer.Each(func(i int, s *goquery.Selection) {
		this.config.parser.removeNode(s)
	})
	return doc
}

func (this *cleaner) cleanAside(doc *goquery.Document) *goquery.Document {
	aside := doc.Find("aside")
	aside.Each(func(i int, s *goquery.Selection) {
		this.config.parser.removeNode(s)
	})
	return doc
}

func (this *cleaner) cleanCites(doc *goquery.Document) *goquery.Document {
	cites := doc.Find("cite")
	cites.Each(func(i int, s *goquery.Selection) {
		this.config.parser.removeNode(s)
	})
	return doc
}

func (this *cleaner) cleanDivs(doc *goquery.Document) *goquery.Document {
	frames := make(map[string]int)
	framesNodes := make(map[string]*list.List)
	divs := doc.Find("div")
	divs.Each(func(i int, s *goquery.Selection) {
		children := s.Children()
		if children.Size() == 0 {
			text := s.Text()
			text = strings.Trim(text, " ")
			text = strings.Trim(text, "\t")
			text = strings.ToLower(text)
			frames[text]++
			if framesNodes[text] == nil {
				framesNodes[text] = list.New()
			}
			framesNodes[text].PushBack(s)
		}
	})
	for text, freq := range frames {
		if freq > 1 {
			selections := framesNodes[text]
			for s := selections.Front(); s != nil; s = s.Next() {
				selection := s.Value.(*goquery.Selection)
				this.config.parser.removeNode(selection)
			}
		}
	}
	return doc
}

func (this *cleaner) dropCaps(doc *goquery.Document) *goquery.Document {
	items := doc.Find("span")
	count := 0 //remove
	items.Each(func(i int, s *goquery.Selection) {
		attribute, exists := s.Attr("class")
		if exists && (strings.Contains(attribute, "dropcap") || strings.Contains(attribute, "drop_cap")) {
			count++
			this.config.parser.dropTag(s)
		}
	})
	if this.config.debug {
		log.Printf("Cleaning %d dropcap tags\n", count)
	}
	return doc
}

func (this *cleaner) removeScriptsStyle(doc *goquery.Document) *goquery.Document {
	if this.config.debug {
		log.Println("Starting to remove script tags")
	}
	scripts := doc.Find("script,noscript,style")
	scripts.Each(func(i int, s *goquery.Selection) {
		this.config.parser.removeNode(s)
	})
	if this.config.debug {
		log.Printf("Removed %d script and style tags\n", scripts.Size())
	}

	//remove comments :) How????
	return doc
}

func (this *cleaner) matchNodeRegEx(attribute string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(attribute)
}

func (this *cleaner) removeNodesRegEx(doc *goquery.Document, pattern *regexp.Regexp) *goquery.Document {
	selectors := [3]string{"id", "class", "name"}
	for _, selector := range selectors {
		naughtyList := doc.Find("*[" + selector + "]")
		cont := 0
		naughtyList.Each(func(i int, s *goquery.Selection) {
			attribute, _ := s.Attr(selector)
			if this.matchNodeRegEx(attribute, pattern) {
				cont++
				this.config.parser.removeNode(s)
			}
		})

		if this.config.debug {
			log.Printf("regExRemoveNodes %d %s elements found against pattern %s\n", cont, selector, pattern.String())
		}
	}
	return doc
}

func (this *cleaner) cleanBadTags(doc *goquery.Document) *goquery.Document {
	body := doc.Find("body")
	children := body.Children()
	selectors := []string{"id", "class", "name"}
	for _, selector := range selectors {
		children.Each(func(i int, s *goquery.Selection) {
			naughtyList := s.Find("*[" + selector + "]")
			cont := 0
			naughtyList.Each(func(j int, e *goquery.Selection) {
				attribute, _ := e.Attr(selector)
				if this.matchNodeRegEx(attribute, REMOVENODES_RE) {
					if this.config.debug {

						log.Printf("Cleaning: Removing node with %s: %s\n", selector, this.config.parser.name(selector, e))
					}
					this.config.parser.removeNode(e)
					cont++
				}
			})
			if this.config.debug && cont > 0 {
				log.Printf("%d naughty %s elements found", cont, selector)
			}
		})
	}
	return doc
}

func (this *cleaner) cleanParaSpans(doc *goquery.Document) *goquery.Document {
	spans := doc.Find("span")
	spans.Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		if parent != nil && parent.Length() > 0 && parent.Get(0).DataAtom == atom.P {
			node := s.Get(0)
			node.Data = s.Text()
			node.Type = html.TextNode
		}
	})
	return doc
}

func (this *cleaner) getFlushedBuffer(fragment string) []*html.Node {
	output := make([]*html.Node, 0)
	reader := strings.NewReader(fragment)
	document, _ := html.Parse(reader)
	body := document.FirstChild.LastChild
	for c := body.FirstChild; c != nil; c = c.NextSibling {
		output = append(output, c)
		c.Parent = nil
		c.PrevSibling = nil
	}

	for _, o := range output {
		o.NextSibling = nil
	}
	return output
}

func (this *cleaner) replaceWithPara(div *goquery.Selection) {
	if div.Size() > 0 {
		node := div.Get(0)
		node.Data = atom.P.String()
		node.DataAtom = atom.P
	}
}

func (this *cleaner) tabsAndNewLinesReplacements(text string) string {
	text = strings.Replace(text, "\n", "\n\n", -1)
	text = tabsRegEx.ReplaceAllString(text, "")
	return text
}

func (this *cleaner) convertDivsToParagraphs(doc *goquery.Document, domType string) *goquery.Document {
	if this.config.debug {
		log.Println("Starting to replace bad divs...")
	}
	badDivs := 0
	convertedTextNodes := 0
	divs := doc.Find(domType)

	divs.Each(func(i int, div *goquery.Selection) {
		divHtml, _ := div.Html()
		if divToPElementsPattern.Match([]byte(divHtml)) {
			this.replaceWithPara(div)
			badDivs++
		} else {
			replacementText := make([]string, 0)
			nodesToRemove := list.New()
			children := div.Contents()
			if this.config.debug {
				log.Printf("Found %d children of div\n", children.Size())
			}
			children.EachWithBreak(func(i int, kid *goquery.Selection) bool {
				text := kid.Text()
				kidNode := kid.Get(0)
				tag := kidNode.Data
				if tag == text {
					tag = "#text"
				}
				if tag == "#text" {
					text = strings.Replace(text, "\n", "", -1)
					text = tabsRegEx.ReplaceAllString(text, "")
					if text == "" {
						return true
					}
					if len(text) > 1 {
						prev := kidNode.PrevSibling
						if this.config.debug {
							log.Printf("PARENT CLASS: %s NODENAME: %s\n", this.config.parser.name("class", div), tag)
							log.Printf("TEXTREPLACE: %s\n", strings.Replace(text, "\n", "", -1))
						}
						if prev != nil && prev.DataAtom == atom.A {
							nodeSelection := kid.HasNodes(prev)
							html, _ := nodeSelection.Html()
							replacementText = append(replacementText, html)
							if this.config.debug {
								log.Printf("SIBLING NODENAME ADDITION: %s TEXT: %s\n", prev.Data, html)
							}
						}
						replacementText = append(replacementText, text)
						nodesToRemove.PushBack(kidNode)
						convertedTextNodes++
					}

				}
				return true
			})

			newNode := new(html.Node)
			newNode.Type = html.ElementNode
			newNode.Data = strings.Join(replacementText, "")
			newNode.DataAtom = atom.P
			div.First().AddNodes(newNode)

			for s := nodesToRemove.Front(); s != nil; s = s.Next() {
				node := s.Value.(*html.Node)
				if node != nil && node.Parent != nil {
					node.Parent.RemoveChild(node)
				}
			}
		}
	})
	if this.config.debug {
		log.Printf("Found %d total divs with %d bad divs replaced and %d textnodes converted inside divs", divs.Size(), badDivs, convertedTextNodes)
	}
	return doc

}
