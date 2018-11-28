package docx_test

import (
	"os"
	"strings"
	"testing"

	"code.sajari.com/docconv"
	_ "code.sajari.com/docconv/docx_test/resources"
)

func TestConvertDocx(t *testing.T) {
	f, err := os.Open("./resources/sample.docx")
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	resp, _, err := docconv.ConvertDocx(f)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}

	if want := "Header"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
	if want := "Footer"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
	if want := "Content"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
}

func TestConvertDocxWithUncommonValidStructure(t *testing.T) {
	f, err := os.Open("./resources/sample_2.docx")
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	resp, _, err := docconv.ConvertDocx(f)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}

	if want := "Header"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
	if want := "Footer"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
	if want := "Content"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
}
