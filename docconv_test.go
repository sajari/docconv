package docconv

import (
	"strings"
	"testing"
)

func TestConvertTrimsSpace(t *testing.T) {
	resp, err := Convert(
		strings.NewReader(" \n\n\nthe \n file\n\n"),
		"text/plain",
		false,
	)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	if want := "the \n file"; resp.Body != want {
		t.Errorf("body = %v, want %v", resp.Body, want)
	}
}
