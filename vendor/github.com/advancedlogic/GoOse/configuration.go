package goose

import "time"

const defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7"

// Configuration is a wrapper for various config options
type Configuration struct {
	localStoragePath        string //not used in this version
	imagesMinBytes          int    //not used in this version
	enableImageFetching     bool
	useMetaLanguage         bool
	targetLanguage          string
	imageMagickConvertPath  string //not used in this version
	imageMagickIdentifyPath string //not used in this version
	browserUserAgent        string
	debug                   bool
	extractPublishDate      bool
	additionalDataExtractor bool

	//path to the stopwords folder
	stopWordsPath string
	stopWords     StopWords
	parser        *Parser

	timeout time.Duration
}

// GetDefaultConfiguration returns safe default configuration options
func GetDefaultConfiguration(args ...string) Configuration {
	if len(args) == 0 {
		return Configuration{
			localStoragePath:        "",   //not used in this version
			imagesMinBytes:          4500, //not used in this version
			enableImageFetching:     true,
			useMetaLanguage:         true,
			targetLanguage:          "en",
			imageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
			imageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
			browserUserAgent:        defaultUserAgent,
			debug:                   false,
			extractPublishDate:      false,
			additionalDataExtractor: false,
			stopWordsPath:           "resources/stopwords",
			stopWords:               NewStopwords(), //TODO with path
			parser:                  NewParser(),
			timeout:                 time.Duration(5 * time.Second),
		}
	}
	return Configuration{
		localStoragePath:        "",   //not used in this version
		imagesMinBytes:          4500, //not used in this version
		enableImageFetching:     true,
		useMetaLanguage:         true,
		targetLanguage:          "en",
		imageMagickConvertPath:  "/usr/bin/convert",  //not used in this version
		imageMagickIdentifyPath: "/usr/bin/identify", //not used in this version
		browserUserAgent:        defaultUserAgent,
		debug:                   false,
		extractPublishDate:      false,
		additionalDataExtractor: false,
		stopWordsPath:           "resources/stopwords",
		stopWords:               NewStopwords(), //TODO with path
		parser:                  NewParser(),
		timeout:                 time.Duration(5 * time.Second),
	}

	/*
		path := args[0]
		jsconfiguration, err := jsconfig.LoadConfig(path)
		if err != nil {
			panic(err.Error())
		}
		stopWordsPath := jsconfiguration.String("stopWordsPath", "resources/stopwords")
		stopWords := NewStopwords() //TODO with path
		return Configuration{
			localStoragePath:        jsconfiguration.String("localStoragePath", ""), //not used in this version
			imagesMinBytes:          jsconfiguration.Int("imageMinBytes", 4500),     //not used in this version
			enableImageFetching:     jsconfiguration.Bool("enableImageFetching", true),
			useMetaLanguage:         jsconfiguration.Bool("useMetaLanguage", true),
			targetLanguage:          jsconfiguration.String("targetLanguage", "en"),
			imageMagickConvertPath:  jsconfiguration.String("imageMagickConvertPath", "/usr/bin/convert"),   //not used in this version
			imageMagickIdentifyPath: jsconfiguration.String("imageMagickIdentityPath", "/usr/bin/identify"), //not used in this version
			browserUserAgent:        jsconfiguration.String("browserUserAgent", defaultUserAgent),
			debug:                   jsconfiguration.Bool("debug", false),
			extractPublishDate:      jsconfiguration.Bool("extractPublishDate", false),      //TODO
			additionalDataExtractor: jsconfiguration.Bool("additionalDataExtractor", false), //TODO
			stopWordsPath:           stopWordsPath,
			stopWords:               stopWords,
			parser:                  NewParser(),
			timeout:                 time.Duration(jsconfiguration.Int("timeout", 5)) * time.Second,
		}
	*/
}
