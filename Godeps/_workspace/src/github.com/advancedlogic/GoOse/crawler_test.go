package goose

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func ReadRawHtml(a Article) string {
	path := fmt.Sprintf("sites/%s.html", a.Domain)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("cannot read %q", path))
	}

	return string(file)
}

func ValidateArticle(a Article) error {
	g := New()
	expected := g.ExtractFromRawHtml(a.FinalUrl, ReadRawHtml(a))

	if expected.Title != a.Title {
		return fmt.Errorf("article title does not match. Got %q", expected.Title)
	}

	if expected.MetaDescription != a.MetaDescription {
		return fmt.Errorf("article metaDescription does not match. Got %q", expected.MetaDescription)
	}

	if !strings.Contains(expected.CleanedText, a.CleanedText) {
		return fmt.Errorf("article cleanedText does not contains %q", a.CleanedText)
	}

	if expected.MetaKeywords != a.MetaKeywords {
		return fmt.Errorf("article keywords does not match. Got %q", expected.MetaKeywords)
	}
	if expected.CanonicalLink != a.CanonicalLink {
		return fmt.Errorf("article CanonicalLink does not match. Got %q", expected.CanonicalLink)
	}

	if expected.TopImage != a.TopImage {
		return fmt.Errorf("article topImage does not match. Got %q", expected.TopImage)
	}

	return nil
}

func Test_GloboEsporteParse(t *testing.T) {
	article := Article{
		Domain:          "globoesporte.globo.com",
		Title:           "Rodrigo Caio treina até nas férias e tenta acelerar retorno aos gramados",
		MetaDescription: "Rodrigo Caio treina na esteira durante as férias em Dracena-SP (Foto: Divulgação)Rodrigo Caio quer ganhar tempo na recuperação da lesão que sofreu no joelho esquerdo. Apesar de ter sido liberado pelo departamento médico do São Paulo para as férias, o ...",
		CleanedText:     "Comissão técnica planeja volta dele para o fim de fevereiro ou início de março",
		MetaKeywords:    "notícias, notícia, presidente prudente região",
		CanonicalLink:   "http://globoesporte.globo.com/sp/presidente-prudente-regiao/noticia/2014/12/rodrigo-caio-treina-ate-nas-ferias-e-tenta-acelerar-retorno-aos-gramados.html",
		TopImage:        "http://s.glbimg.com/es/ge/f/original/2014/12/26/10863872_894379987249341_2406060334390226774_o.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

func Test_EditionCnnParse(t *testing.T) {
	article := Article{
		Domain:          "edition.cnn.com",
		Title:           "What if you could make anything you wanted?",
		MetaDescription: "Massimo Banzi's pocket-sized open-source circuit board has become a key building block in the creation of a huge variety of innovative devices.",
		CleanedText:     "In the 20th century, getting your child a toy car meant a trip to a shopping mall.",
		MetaKeywords:    "",
		CanonicalLink:   "http://www.cnn.com/2012/07/08/opinion/banzi-ted-open-source/index.html",
		TopImage:        "http://i2.cdn.turner.com/cnn/dam/assets/120706022111-ted-cnn-ideas-massimo-banzi-00003302-story-top.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}

func Test_BbcParse(t *testing.T) {
	article := Article{
		Domain:          "bbc.com",
		Title:           "Crunch talks on new Greek bailout",
		MetaDescription: "German and Greek finance ministers meet IMF and Eurogroup chiefs ahead of a crucial finance ministers' meeting on Greece's bailout request.",
		CleanedText:     "Mr Tsipras won elections in late January on a platform of rejecting the austerity measures tied to the bailout.",
		MetaKeywords:    "keywords, added, to, test, case insensitive",
		CanonicalLink:   "http://www.bbc.com/news/business-31545115",
		TopImage:        "http://news.bbcimg.co.uk/media/images/81120000/jpg/_81120901_81120501.jpg",
	}

	err := ValidateArticle(article)
	if err != nil {
		t.Error(err)
	}
}
