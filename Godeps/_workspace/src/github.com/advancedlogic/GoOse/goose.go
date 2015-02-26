package goose

type Goose struct {
	config configuration
}

func New(args ...string) Goose {

	return Goose{
		config: GetDefualtConfiguration(args...),
	}
}

func (this Goose) ExtractFromUrl(url string) *Article {
	cc := NewCrawler(this.config, url, "")
	return cc.Crawl()
}

func (this Goose) ExtractFromRawHtml(url string, rawHtml string) *Article {
	cc := NewCrawler(this.config, url, rawHtml)
	return cc.Crawl()
}
