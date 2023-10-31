package docx_test

import (
	"os"
	"strings"
	"testing"

	"code.sajari.com/docconv/v2"
)

func TestConvertPptx(t *testing.T) {
	f, err := os.Open("./testdata/sample.pptx")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}

	resp, _, err := docconv.ConvertPptx(f)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	if want := "Get text from pptx"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contain %v", resp, want)
	}
	if want := "First"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contain %v", resp, want)
	}
}
