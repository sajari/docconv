package goose

import (
	"gopkg.in/fatih/set.v0"
)

//some word statistics
type wordStats struct {
	//total number of stopwords or good words that we can calculate
	stopWordCount int
	//total number of words on a node
	wordCount int
	//holds an actual list of the stop words we found
	stopWords *set.Set
}

func (this *wordStats) getStopWords() *set.Set {
	return this.stopWords
}

func (this *wordStats) setStopWords(stopWords *set.Set) {
	this.stopWords = stopWords
}

func (this *wordStats) getStopWordCount() int {
	return this.stopWordCount
}

func (this *wordStats) setStopWordCount(stopWordCount int) {
	this.stopWordCount = stopWordCount
}

func (this *wordStats) getWordCount() int {
	return this.wordCount
}

func (this *wordStats) setWordCount(wordCount int) {
	this.wordCount = wordCount
}
