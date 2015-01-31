package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Convert RTF
func ConvertRtf(input io.Reader) (string, map[string]string) {

	// Save input to a file
	inputFile, err := ioutil.TempFile("/tmp", "sajari-convert-")
	if err != nil {
		log.Println("TempFile:", err)
	}
	defer os.Remove(inputFile.Name())
	io.Copy(inputFile, input)

	var output string
	tmpOutput, err := exec.Command("unrtf", "--nopict", "--text", inputFile.Name()).Output()
	if err != nil {
		log.Println("unrtf:", err)
	}

	// Step through content looking for meta data and stripping out comments
	info := make(map[string]string)
	for _, line := range strings.Split(string(tmpOutput), "\n") {
		if parts := strings.SplitN(line, ":", 2); len(parts) > 1 {
			info[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
		if len(line) > 4 && line[:4] != "### " {
			output += line + "\n"
		}
	}

	// Identify meta data
	meta := make(map[string]string)
	if tmp, ok := info["AUTHOR"]; ok {
		meta["Author"] = tmp
	}
	if tmp, ok := info["### creation date"]; ok {
		if t, err := time.Parse("02 January 2006 15:04", tmp); err == nil {
			meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
		}
	}
	if tmp, ok := info["### revision date"]; ok {
		if t, err := time.Parse("02 January 2006 15:04", tmp); err == nil {
			meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
		}
	}

	return output, meta
}
