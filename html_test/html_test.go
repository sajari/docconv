package html_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"code.sajari.com/docconv"
)

func TestConvertHTML(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.html")
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	res, _, err := docconv.ConvertHTML(bytes.NewReader(data), true)
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
	lines := strings.Split(res, "\n")
	for _, i := range lines {
		t.Log(i)
	}
	line, expectedLine := lines[0], "1"
	if line != expectedLine {
		t.Fatalf("got %s, want %s", line, expectedLine)
	}
	line, expectedLine = lines[1], "word"
	if line != expectedLine {
		t.Fatalf("got %s, want %s", line, expectedLine)
	}
	line, expectedLine = lines[2], "This is a full sentence."
	if line != expectedLine {
		t.Fatalf("got %s, want %s", line, expectedLine)
	}
}
