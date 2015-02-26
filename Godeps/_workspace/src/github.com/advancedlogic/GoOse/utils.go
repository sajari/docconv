package goose

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

func ReadLinesOfFile(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func TimeInMilliseconds() int64 {
	now := time.Now()
	return now.Unix()
}

func TimeInNanoseconds() int64 {
	now := time.Now()
	return now.UnixNano()
}

func RegSplit(text string, reg *regexp.Regexp) []string {
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}
