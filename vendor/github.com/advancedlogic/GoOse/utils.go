package goose

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// ReadLinesOfFile returns the lines from a file as a slice of strings
func ReadLinesOfFile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

// RegSplit splits the strings into strings using the regular expression as separator
func RegSplit(text string, reg *regexp.Regexp) []string {
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:]
	return result
}
