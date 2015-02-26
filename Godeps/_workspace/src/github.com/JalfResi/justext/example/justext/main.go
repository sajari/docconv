package main

import (
	"flag"
	"github.com/JalfResi/gojustext"
	"log"
	"os"
)

var (
	lengthLow          *int     = flag.Int("length-low", 70, "Sets the length low")
	lengthHigh         *int     = flag.Int("length-high", 200, "Sets the length high")
	stoplistFilename   *string  = flag.String("stoplist", "", "Specifies a stoplist file")
	stoplistLanguage   *string  = flag.String("language", "English", "Sets the stoplist language")
	stopwordsLow       *float64 = flag.Float64("stopwords-low", 0.30, "Sets the stopwords low")
	stopwordsHigh      *float64 = flag.Float64("stopwords-high", 0.32, "Sets the stopwords high")
	maxLinkDensity     *float64 = flag.Float64("max-link-density", 0.2, "Sets the max link density")
	maxHeadingDistance *int     = flag.Int("max-heading-distance", 200, "Sets the max heading distance")
	noHeadings         *bool    = flag.Bool("no-headings", false, "Ignores headings in output")
	outputMode         *bool    = flag.Bool("detailed", false, "Generates HTML with detailed results")
	debugMode          *bool    = flag.Bool("debug", false, "Log debug messages")
)

func main() {
	flag.Parse()

	var stoplist map[string]bool

	if *stoplistFilename != "" {
		stoplist = justext.ReadStoplist(*stoplistFilename)
	} else {
		var err error
		stoplist, err = justext.GetStoplist(*stoplistLanguage)
		if err != nil {
			log.Fatal(err)
		}
	}

	jusText := justext.NewReader(os.Stdin)

	jusText.LengthLow = *lengthLow
	jusText.LengthHigh = *lengthHigh
	jusText.Stoplist = stoplist
	jusText.StopwordsLow = *stopwordsLow
	jusText.StopwordsHigh = *stopwordsHigh
	jusText.MaxLinkDensity = *maxLinkDensity
	jusText.MaxHeadingDistance = *maxHeadingDistance
	jusText.NoHeadings = *noHeadings

	p, err := jusText.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	j := justext.NewWriter(os.Stdout)
	j.Stoplist = stoplist

	if *outputMode {
		j.Mode = justext.MODE_DETAILED
	}

	if *debugMode {
		j.OutputDebug(p)
	}

	j.WriteAll(p)
}
