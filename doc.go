package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Convert MS Word DOC
func ConvertDoc(input io.Reader) (string, map[string]string) {

	// Save input to a file
	inputFile, err := ioutil.TempFile("/tmp", "sajari-convert-")
	if err != nil {
		log.Println("TempFile:", err)
	}
	defer os.Remove(inputFile.Name())
	io.Copy(inputFile, input)

	// Meta data
	mc := make(chan map[string]string, 1)
	go func() {
		meta := make(map[string]string)
		metaStr, err := exec.Command("wvSummary", inputFile.Name()).Output()
		if err != nil {
			log.Println("wvSummary:", err)
		}

		// Parse meta output
		info := make(map[string]string)
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, "=", 2); len(parts) > 1 {
				info[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed meta
		if tmp, ok := info["Last Modified"]; ok {
			if t, err := time.Parse(time.RFC3339, tmp); err == nil {
				meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if tmp, ok := info["Created"]; ok {
			if t, err := time.Parse(time.RFC3339, tmp); err == nil {
				meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}

		mc <- meta
	}()

	// Document body
	bc := make(chan string, 1)
	go func() {

		// Save output to a file
		outputFile, err := ioutil.TempFile("/tmp", "sajari-convert-")
		if err != nil {
			log.Println("TempFile Out:", err)
			return
		}
		defer os.Remove(outputFile.Name())

		err = exec.Command("wvText", inputFile.Name(), outputFile.Name()).Run()
		if err != nil {
			log.Println("wvText:", err)
		}

		var buf bytes.Buffer
		_, err = buf.ReadFrom(outputFile)
		if err != nil {
			log.Println("wvText:", err)
		}

		bc <- buf.String()
	}()

	body := <-bc
	meta := <-mc

	if len(body) == 0 {
		inputFile.Seek(0, 0)
		return ConvertDocx(inputFile)
	}
	return body, meta
}
