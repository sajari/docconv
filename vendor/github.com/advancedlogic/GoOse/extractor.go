package goose

import (
	"container/list"
	"github.com/advancedlogic/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gopkg.in/fatih/set.v0"
	"log"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const defaultLanguage = "en"

var motleyReplacement = "&#65533;"
var escapedFragmentReplacement = regexp.MustCompile("#!")
var titleReplacements = regexp.MustCompile("&raquo;")

var pipeSplitter = regexp.MustCompile("\\|")
var dashSplitter = regexp.MustCompile(" - ")
var arrowsSplitter = regexp.MustCompile("»")
var colonSplitter = regexp.MustCompile(":")
var spaceSplitter = regexp.MustCompile(" ")

var aRelTagSelector = "a[rel=tag]"
var aHrefTagSelector = [...]string{"/tag/", "/tags/", "/topic/", "?keyword"}
var langRegEx = "^[A-Za-z]{2}$"

// ContentExtractor can parse the HTML and fetch various properties
type ContentExtractor struct {
	config Configuration
}

// NewExtractor returns a configured HTML parser
func NewExtractor(config Configuration) ContentExtractor {
	return ContentExtractor{
		config: config,
	}
}

//if the article has a title set in the source, use that
func (extr *ContentExtractor) getTitle(article *Article) string {
	title := ""
	doc := article.Doc

	ogTitleElement := doc.Find(`meta[property="og:title"]`)
	if ogTitleElement != nil && ogTitleElement.Size() > 0 {
		title, _ = ogTitleElement.Attr("content")
	}
	if title == "" {
		titleElement := doc.Find("title,post-title,headline")
		if titleElement == nil || titleElement.Size() == 0 {
			return title
		}
		title = titleElement.Text()
	}

	usedDelimiter := false

	if strings.Contains(title, "|") {
		title = extr.splitTitle(RegSplit(title, pipeSplitter))
		usedDelimiter = true
	}

	if !usedDelimiter && strings.Contains(title, "-") {
		title = extr.splitTitle(RegSplit(title, dashSplitter))
		usedDelimiter = true
	}

	if !usedDelimiter && strings.Contains(title, "»") {
		title = extr.splitTitle(RegSplit(title, arrowsSplitter))
		usedDelimiter = true
	}

	if !usedDelimiter && strings.Contains(title, ":") {
		title = extr.splitTitle(RegSplit(title, colonSplitter))
		usedDelimiter = true
	}

	title = strings.Replace(title, motleyReplacement, "", -1)

	if extr.config.debug {
		log.Printf("Page title is %s\n", title)
	}

	return strings.TrimSpace(title)
}

func (extr *ContentExtractor) splitTitle(titles []string) string {
	largeTextLength := 0
	largeTextIndex := 0
	for i, current := range titles {
		if len(current) > largeTextLength {
			largeTextLength = len(current)
			largeTextIndex = i
		}
	}
	title := titles[largeTextIndex]
	title = strings.Replace(title, "&raquo;", "»", -1)
	return title
}

//if the article has meta language set in the source, use that
func (extr *ContentExtractor) getMetaLanguage(article *Article) string {
	language := ""
	doc := article.Doc
	shtml := doc.Find("html")
	attr, _ := shtml.Attr("lang")
	if attr == "" {
		attr, _ = doc.Attr("lang")
	}
	if attr == "" {
		selection := doc.Find("meta").EachWithBreak(func(i int, s *goquery.Selection) bool {
			exists := false
			attr, exists = s.Attr("http-equiv")
			if exists && attr == "content-language" {
				return false
			}
			return true
		})
		if selection != nil {
			attr, _ = selection.Attr("content")
		}
	}
	idx := strings.LastIndex(attr, "-")
	if idx == -1 {
		language = attr
	} else {
		language = attr[0:idx]
	}

	_, ok := sw[language]

	if language == "" || !ok {
		language = extr.config.stopWords.SimpleLanguageDetector(shtml.Text())
		if language == "" {
			language = defaultLanguage
		}
	}

	extr.config.targetLanguage = language
	return language
}

//if the article has favicon set in the source, use that
func (extr *ContentExtractor) getFavicon(article *Article) string {
	favicon := ""
	doc := article.Doc
	doc.Find("link").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attr, exists := s.Attr("rel")
		if exists && strings.Contains(attr, "icon") {
			favicon, _ = s.Attr("href")
			return false
		}
		return true
	})
	return favicon
}

func (extr *ContentExtractor) getMetaContentWithSelector(article *Article, selector string) string {
	content := ""
	doc := article.Doc
	selection := doc.Find(selector)
	content, _ = selection.Attr("content")
	return strings.TrimSpace(content)
}

func (extr *ContentExtractor) getMetaContent(article *Article, metaName string) string {
	content := ""
	doc := article.Doc
	doc.Find("meta").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attr, exists := s.Attr("name")
		if exists && attr == metaName {
			content, _ = s.Attr("content")
			return false
		}
		attr, exists = s.Attr("itemprop")
		if exists && attr == metaName {
			content, _ = s.Attr("content")
			return false
		}
		return true
	})
	return content
}

func (extr *ContentExtractor) getMetaContents(article *Article, metaNames *set.Set) map[string]string {
	contents := make(map[string]string)
	doc := article.Doc
	counter := metaNames.Size()
	doc.Find("meta").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attr, exists := s.Attr("name")
		if exists && metaNames.Has(attr) {
			content, _ := s.Attr("content")
			contents[attr] = content
			counter--
			if counter < 0 {
				return false
			}
		}
		return true
	})
	return contents
}

//if the article has meta description set in the source, use that
func (extr *ContentExtractor) getMetaDescription(article *Article) string {
	return extr.getMetaContent(article, "description")
}

//if the article has meta keywords set in the source, use that
func (extr *ContentExtractor) getMetaKeywords(article *Article) string {
	return extr.getMetaContent(article, "keywords")
}

//if the article has meta author set in the source, use that
func (extr *ContentExtractor) getMetaAuthor(article *Article) string {
	return extr.getMetaContent(article, "author")
}

//if the article has meta content location set in the source, use that
func (extr *ContentExtractor) getMetaContentLocation(article *Article) string {
	return extr.getMetaContent(article, "contentLocation")
}

//if the article has meta canonical link set in the url
func (extr *ContentExtractor) getCanonicalLink(article *Article) string {
	doc := article.Doc
	metas := doc.Find("link[rel=canonical]")
	if metas.Length() > 0 {
		meta := metas.First()
		href, _ := meta.Attr("href")
		href = strings.Trim(href, "\n")
		href = strings.Trim(href, " ")
		if href != "" {
			return href
		}
	}
	return article.FinalURL
}

//extract domain and use that
func (extr *ContentExtractor) getDomain(article *Article) string {
	canonicalLink := article.CanonicalLink
	u, err := url.Parse(canonicalLink)
	if err == nil {
		return u.Host
	}
	return ""
}

//if the article has tags set in the source, use that
func (extr *ContentExtractor) getTags(article *Article) *set.Set {
	tags := set.New()
	doc := article.Doc
	selections := doc.Find(aRelTagSelector)
	selections.Each(func(i int, s *goquery.Selection) {
		tags.Add(s.Text())
	})
	selections = doc.Find("a")
	selections.Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			for _, part := range aHrefTagSelector {
				if strings.Contains(href, part) {
					tags.Add(s.Text())
				}
			}
		}
	})

	return tags
}

//we're going to start looking for where the clusters of paragraphs are. We'll score a cluster based on the number of stopwords
//and the number of consecutive paragraphs together, which should form the cluster of text that this node is around
//also store on how high up the paragraphs are, comments are usually at the bottom and should get a lower score
func (extr *ContentExtractor) calculateBestNode(article *Article) *goquery.Selection {
	doc := article.Doc
	var topNode *goquery.Selection
	nodesToCheck := extr.nodesToCheck(doc)
	if extr.config.debug {
		log.Printf("Nodes to check %d\n", len(nodesToCheck))
	}
	startingBoost := 1.0
	cnt := 0
	i := 0
	parentNodes := set.New()
	nodesWithText := list.New()
	for _, node := range nodesToCheck {
		textNode := node.Text()
		ws := extr.config.stopWords.stopWordsCount(extr.config.targetLanguage, textNode)
		highLinkDensity := extr.isHighLinkDensity(node)
		if ws.stopWordCount > 2 && !highLinkDensity {
			nodesWithText.PushBack(node)
		}
	}
	nodesNumber := nodesWithText.Len()
	negativeScoring := 0
	bottomNegativeScoring := float64(nodesNumber) * 0.25

	if extr.config.debug {
		log.Printf("About to inspect num of nodes with text %d\n", nodesNumber)
	}

	for n := nodesWithText.Front(); n != nil; n = n.Next() {
		node := n.Value.(*goquery.Selection)
		boostScore := 0.0
		if extr.isBoostable(node) {
			if cnt >= 0 {
				boostScore = float64((1.0 / startingBoost) * 50)
				startingBoost++
			}
		}

		if nodesNumber > 15 {
			if float64(nodesNumber-i) <= bottomNegativeScoring {
				booster := bottomNegativeScoring - float64(nodesNumber-i)
				boostScore = -math.Pow(booster, 2.0)
				negScore := math.Abs(boostScore) + float64(negativeScoring)
				if negScore > 40 {
					boostScore = 5.0
				}
			}
		}

		if extr.config.debug {
			log.Printf("Location Boost Score %1.5f on iteration %d id='%s' class='%s'\n", boostScore, i, extr.config.parser.name("id", node), extr.config.parser.name("class", node))
		}
		textNode := node.Text()
		ws := extr.config.stopWords.stopWordsCount(extr.config.targetLanguage, textNode)
		upScore := ws.stopWordCount + int(boostScore)
		parentNode := node.Parent()
		extr.updateScore(parentNode, upScore)
		extr.updateNodeCount(parentNode, 1)
		if !parentNodes.Has(parentNode) {
			parentNodes.Add(parentNode)
		}
		parentParentNode := parentNode.Parent()
		if parentParentNode != nil {
			extr.updateNodeCount(parentParentNode, 1)
			extr.updateScore(parentParentNode, upScore/2.0)
			if !parentNodes.Has(parentParentNode) {
				parentNodes.Add(parentParentNode)
			}
		}
		cnt++
		i++
	}

	topNodeScore := 0
	parentNodesArray := parentNodes.List()
	for _, p := range parentNodesArray {
		e := p.(*goquery.Selection)
		if extr.config.debug {
			log.Printf("ParentNode: score=%s nodeCount=%s id='%s' class='%s'\n", extr.config.parser.name("gravityScore", e), extr.config.parser.name("gravityNodes", e), extr.config.parser.name("id", e), extr.config.parser.name("class", e))
		}
		score := extr.getScore(e)
		if score >= topNodeScore {
			topNode = e
			topNodeScore = score
		}
		if topNode == nil {
			topNode = e
		}
	}
	return topNode
}

//returns the gravityScore as an integer from this node
func (extr *ContentExtractor) getScore(node *goquery.Selection) int {
	return extr.getNodeGravityScore(node)
}

func (extr *ContentExtractor) getNodeGravityScore(node *goquery.Selection) int {
	grvScoreString, exists := node.Attr("gravityScore")
	if !exists {
		return 0
	}
	grvScore, err := strconv.Atoi(grvScoreString)
	if err != nil {
		return 0
	}
	return grvScore
}

//adds a score to the gravityScore Attribute we put on divs
//we'll get the current score then add the score we're passing in to the current
func (extr *ContentExtractor) updateScore(node *goquery.Selection, addToScore int) {
	currentScore := 0
	var err error
	scoreString, _ := node.Attr("gravityScore")
	if scoreString != "" {
		currentScore, err = strconv.Atoi(scoreString)
		if err != nil {
			currentScore = 0
		}
	}
	newScore := currentScore + addToScore
	extr.config.parser.setAttr(node, "gravityScore", strconv.Itoa(newScore))
}

//stores how many decent nodes are under a parent node
func (extr *ContentExtractor) updateNodeCount(node *goquery.Selection, addToCount int) {
	currentScore := 0
	var err error
	scoreString, _ := node.Attr("gravityNodes")
	if scoreString != "" {
		currentScore, err = strconv.Atoi(scoreString)
		if err != nil {
			currentScore = 0
		}
	}
	newScore := currentScore + addToCount
	extr.config.parser.setAttr(node, "gravityNodes", strconv.Itoa(newScore))
}

//a lot of times the first paragraph might be the caption under an image so we'll want to make sure if we're going to
//boost a parent node that it should be connected to other paragraphs, at least for the first n paragraphs
//so we'll want to make sure that the next sibling is a paragraph and has at least some substatial weight to it
func (extr *ContentExtractor) isBoostable(node *goquery.Selection) bool {
	stepsAway := 0
	next := node.Next()
	for next != nil && stepsAway < node.Siblings().Length() {
		currentNodeTag := node.Get(0).DataAtom.String()
		if currentNodeTag == "p" {
			if stepsAway >= 3 {
				if extr.config.debug {
					log.Println("Next paragraph is too far away, not boosting")
				}
				return false
			}

			paraText := node.Text()
			ws := extr.config.stopWords.stopWordsCount(extr.config.targetLanguage, paraText)
			if ws.stopWordCount > 5 {
				if extr.config.debug {
					log.Println("We're gonna boost this node, seems content")
				}
				return true
			}
		}

		stepsAway++
		next = next.Next()
	}

	return false
}

//returns a list of nodes we want to search on like paragraphs and tables
func (extr *ContentExtractor) nodesToCheck(doc *goquery.Document) []*goquery.Selection {
	var output []*goquery.Selection
	tags := []string{"p", "pre", "td"}
	for _, tag := range tags {
		selections := doc.Children().Find(tag)
		if selections != nil {
			selections.Each(func(i int, s *goquery.Selection) {
				output = append(output, s)
			})
		}
	}
	return output
}

//checks the density of links within a node, is there not much text and most of it contains bad links?
//if so it's no good
func (extr *ContentExtractor) isHighLinkDensity(node *goquery.Selection) bool {
	links := node.Find("a")
	if links == nil || links.Size() == 0 {
		return false
	}
	text := node.Text()
	words := strings.Split(text, " ")
	nwords := len(words)
	var sb []string
	links.Each(func(i int, s *goquery.Selection) {
		linkText := s.Text()
		sb = append(sb, linkText)
	})
	linkText := strings.Join(sb, "")
	linkWords := strings.Split(linkText, " ")
	nlinkWords := len(linkWords)
	nlinks := links.Size()
	linkDivisor := float64(nlinkWords) / float64(nwords)
	score := linkDivisor * float64(nlinks)

	if extr.config.debug {
		logText := ""
		if len(node.Text()) >= 51 {
			logText = node.Text()[0:50]
		} else {
			logText = node.Text()
		}
		log.Printf("Calculated link density score as %1.5f for node %s\n", score, logText)
	}
	if score > 1.0 {
		return true
	}
	return false
}

func (extr *ContentExtractor) isTableAndNoParaExist(selection *goquery.Selection) bool {
	subParagraph := selection.Find("p")
	subParagraph.Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		if len(txt) < 25 {
			node := s.Get(0)
			parent := node.Parent
			parent.RemoveChild(node)
		}
	})

	subParagraph2 := selection.Find("p")
	if subParagraph2.Length() == 0 && selection.Get(0).DataAtom.String() != "td" {
		return true
	}
	return false
}

func (extr *ContentExtractor) isNodescoreThresholdMet(node *goquery.Selection, e *goquery.Selection) bool {
	topNodeScore := extr.getNodeGravityScore(node)
	currentNodeScore := extr.getNodeGravityScore(e)
	threasholdScore := float64(topNodeScore) * 0.08
	if (float64(currentNodeScore) < threasholdScore) && e.Get(0).DataAtom.String() != "td" {
		return false
	}
	return true
}

//we could have long articles that have tons of paragraphs so if we tried to calculate the base score against
//the total text score of those paragraphs it would be unfair. So we need to normalize the score based on the average scoring
//of the paragraphs within the top node. For example if our total score of 10 paragraphs was 1000 but each had an average value of
//100 then 100 should be our base.
func (extr *ContentExtractor) getSiblingsScore(topNode *goquery.Selection) int {
	base := 100000
	paragraphNumber := 0
	paragraphScore := 0
	nodesToCheck := topNode.Find("p")
	nodesToCheck.Each(func(i int, s *goquery.Selection) {
		textNode := s.Text()
		ws := extr.config.stopWords.stopWordsCount(extr.config.targetLanguage, textNode)
		highLinkDensity := extr.isHighLinkDensity(s)
		if ws.stopWordCount > 2 && !highLinkDensity {
			paragraphNumber++
			paragraphScore += ws.stopWordCount
		}
	})
	if paragraphNumber > 0 {
		base = paragraphScore / paragraphNumber
	}
	return base
}

func (extr *ContentExtractor) getSiblingsContent(currentSibling *goquery.Selection, baselinescoreSiblingsPara float64) []*goquery.Selection {
	var ps []*goquery.Selection
	if currentSibling.Get(0).DataAtom.String() == "p" && len(currentSibling.Text()) > 0 {
		ps = append(ps, currentSibling)
		return ps
	}

	potentialParagraphs := currentSibling.Find("p")
	potentialParagraphs.Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if len(text) > 0 {
			ws := extr.config.stopWords.stopWordsCount(extr.config.targetLanguage, text)
			paragraphScore := ws.stopWordCount
			siblingBaselineScore := 0.30
			highLinkDensity := extr.isHighLinkDensity(s)
			score := siblingBaselineScore * baselinescoreSiblingsPara
			if score < float64(paragraphScore) && !highLinkDensity {
				node := new(html.Node)
				node.Type = html.TextNode
				node.Data = text
				node.DataAtom = atom.P
				nodes := make([]*html.Node, 1)
				nodes[0] = node
				newSelection := new(goquery.Selection)
				newSelection.Nodes = nodes
				ps = append(ps, newSelection)
			}
		}

	})
	return ps
}

func (extr *ContentExtractor) walkSiblings(node *goquery.Selection) []*goquery.Selection {
	currentSibling := node.Prev()
	var b []*goquery.Selection
	for currentSibling.Length() != 0 {
		b = append(b, currentSibling)
		previousSibling := currentSibling.Prev()
		currentSibling = previousSibling
	}
	return b
}

//adds any siblings that may have a decent score to this node
func (extr *ContentExtractor) addSiblings(topNode *goquery.Selection) *goquery.Selection {
	if extr.config.debug {
		log.Println("Starting to add siblings")
	}
	baselinescoreSiblingsPara := extr.getSiblingsScore(topNode)
	results := extr.walkSiblings(topNode)
	for _, currentNode := range results {
		ps := extr.getSiblingsContent(currentNode, float64(baselinescoreSiblingsPara))
		for _, p := range ps {
			nodes := make([]*html.Node, len(topNode.Nodes)+1)
			nodes[0] = p.Get(0)
			for i, node := range topNode.Nodes {
				nodes[i+1] = node
			}
			topNode.Nodes = nodes
		}
	}
	return topNode
}

//remove any divs that looks like non-content, clusters of links, or paras with no gusto
func (extr *ContentExtractor) postCleanup(targetNode *goquery.Selection) *goquery.Selection {
	if extr.config.debug {
		log.Println("Starting cleanup Node")
	}
	node := extr.addSiblings(targetNode)
	children := node.Children()
	children.Each(func(i int, s *goquery.Selection) {
		tag := s.Get(0).DataAtom.String()
		if tag != "p" {
			if extr.config.debug {
				log.Printf("CLEANUP  NODE: %s class: %s\n", extr.config.parser.name("id", s), extr.config.parser.name("class", s))
			}
			//if extr.isHighLinkDensity(s) || extr.isTableAndNoParaExist(s) || !extr.isNodescoreThresholdMet(node, s) {
			if extr.isHighLinkDensity(s) {
				extr.config.parser.removeNode(s)
				return
			}

			subParagraph := s.Find("p")
			subParagraph.Each(func(j int, e *goquery.Selection) {
				if len(e.Text()) < 25 {
					extr.config.parser.removeNode(e)
				}
			})

			subParagraph2 := s.Find("p")
			if subParagraph2.Length() == 0 && tag != "td" {
				if extr.config.debug {
					log.Println("Removing node because it doesn't have any paragraphs")
				}
				extr.config.parser.removeNode(s)
			} else {
				if extr.config.debug {
					log.Println("Not removing TD node")
				}
			}
			return
		}
	})
	return node
}
