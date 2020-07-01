package docconv

// Config settings for doc conv
type Config struct {
	Limitation LenthLimitation
}

// LenthLimitation page or word limit
type LenthLimitation struct {
	// XMLMaxWord max word limit for xml parsing, this will effort office 2007 zip format document
	XMLMaxWord int
	// PdfFirstPage first page to convert for pdf
	PdfFirstPage int
	// PdfLastPage last page to convert for pdf
	PdfLastPage int
}

var (
	config Config
)

// SetConfig set configuration for docconv
func SetConfig(c Config) {
	config = c
}

func checkXMLMaxWord() bool {
	return config.Limitation.XMLMaxWord > 0
}

func xmlMaxWordExceed(length int) bool {
	return length > config.Limitation.XMLMaxWord
}
