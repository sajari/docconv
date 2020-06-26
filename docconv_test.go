package docconv

import (
	"os"
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

func TestXMLMaxWord(t *testing.T) {
	t.Run("max word not set", func(t *testing.T) {
		checkMaxWord := checkXMLMaxWord()
		if checkMaxWord != false {
			t.Fatalf("got %v, want false", checkMaxWord)
		}
	})
	t.Run("test checkMaxWord", func(t *testing.T) {
		SetConfig(Config{Limitation: LenthLimitation{XMLMaxWord: 10}})
		checkMaxWord := checkXMLMaxWord()
		if checkMaxWord != true {
			t.Fatalf("got %v, want true", checkMaxWord)
		}
	})
	t.Run("test xmlMaxWordExceed", func(t *testing.T) {
		SetConfig(Config{Limitation: LenthLimitation{XMLMaxWord: 10}})
		exceed := xmlMaxWordExceed(10)
		if exceed != false {
			t.Fatalf("got %v, want false", exceed)
		}
		exceed = xmlMaxWordExceed(11)
		if exceed != true {
			t.Fatalf("got %v, want true", exceed)
		}
	})
	t.Run("test parse pptx with maxword", func(t *testing.T) {
		SetConfig(Config{Limitation: LenthLimitation{XMLMaxWord: 2}})
		f, err := os.Open("./docx_test/testdata/sample_3.docx")
		if err != nil {
			t.Fatalf("got error = %v, want nil", err)
		}

		resp, _, err := ConvertDocx(f)
		if err != nil {
			t.Fatalf("got error = %v, want nil", err)
		}
		if want := "Content from docx file"; !strings.Contains(resp, want) {
			t.Errorf("expected %v to contains %v", resp, want)
		}
		if want := "second"; strings.Contains(resp, want) {
			t.Errorf("expected %v to not contains %v", resp, want)
		}
	})

}

func TestPDFPageLimit(t *testing.T) {
	SetConfig(Config{Limitation: LenthLimitation{PdfFirstPage: 2, PdfLastPage: 3}})
	f, err := os.Open("./pdf_test/testdata/pdf.pdf")
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}

	resp, _, err := ConvertPDF(f)
	if err != nil {
		t.Fatalf("got error = %v, want nil", err)
	}
	if want := "2"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
	if want := "3"; !strings.Contains(resp, want) {
		t.Errorf("expected %v to contains %v", resp, want)
	}
	if want := "1"; strings.Contains(resp, want) {
		t.Errorf("expected %v to not contains %v", resp, want)
	}
	if want := "4"; strings.Contains(resp, want) {
		t.Errorf("expected %v to not contains %v", resp, want)
	}
	if want := "5"; strings.Contains(resp, want) {
		t.Errorf("expected %v to not contains %v", resp, want)
	}
}
