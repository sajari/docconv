package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Errors & warnings are deliberately suppressed as tidy throws warnings very easily
func Tidy(input io.Reader, xmlIn bool) ([]byte, error) {

	// Save input to a file
	inputFile, err := ioutil.TempFile("/tmp", "sajari-convert-")
	if err != nil {
		log.Println("TempFile:", err)
		return nil, err
	}
	defer os.Remove(inputFile.Name())
	io.Copy(inputFile, input)

	var output []byte
	if xmlIn {
		output, err = exec.Command("tidy", "-xml", "-numeric", "-asxml", "-quiet", "-utf8", inputFile.Name()).Output()
	} else {
		output, err = exec.Command("tidy", "-numeric", "-asxml", "-quiet", "-utf8", inputFile.Name()).Output()
	}

	if err != nil && err.Error() != "exit status 1" {
		return nil, err
	}
	return output, nil
}
