package docconv

import (
	"os"
	"reflect"
	"testing"
)

func TestConvertPDF(t *testing.T) {
	f, err := os.Open("testdata/pdf-test.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	_, meta, err := ConvertPDF(f)
	want := map[string]string{
		"Author":         "Yukon Department of Education",
		"Encrypted":      "yes (print:yes copy:no change:no addNotes:no algorithm:RC4)",
		"ArtBox":         "0.00     0.00   612.00   792.00",
		"File size":      "20597 bytes",
		"PDF version":    "1.6",
		"Creator":        "Acrobat PDFMaker 7.0.7 for Word",
		"ModDate":        "Wed Jun  4 19:47:36 2008 -03",
		"Tagged":         "yes",
		"Suspects":       "no",
		"CropBox":        "0.00     0.00   612.00   792.00",
		"Optimized":      "yes",
		"CreationDate":   "Wed Jun  4 19:44:00 2008 -03",
		"Form":           "none",
		"Page size":      "612 x 792 pts (letter)",
		"Page rot":       "0",
		"TrimBox":        "0.00     0.00   612.00   792.00",
		"MediaBox":       "0.00     0.00   612.00   792.00",
		"BleedBox":       "0.00     0.00   612.00   792.00",
		"Title":          "PDF Test Page",
		"Producer":       "Acrobat Distiller 7.0.5 (Windows)",
		"UserProperties": "no",
		"JavaScript":     "no",
		"Pages":          "1",
	}

	want_v018 := map[string]string{
		"Creator":      "Acrobat PDFMaker 7.0.7 for Word",
		"Producer":     "Acrobat Distiller 7.0.5 (Windows)",
		"CreationDate": "Wed Jun  4 15:44:00 2008",
		"Page size":    "612 x 792 pts (letter)",
		"File size":    "20597 bytes",
		"PDF version":  "1.6", "Title": "PDF Test Page",
		"Author":    "Yukon Department of Education",
		"ModDate":   "Wed Jun  4 15:47:36 2008",
		"Tagged":    "yes",
		"BleedBox":  "0.00     0.00   612.00   792.00",
		"Encrypted": "yes (print:yes copy:no change:no addNotes:no)",
		"MediaBox":  "0.00     0.00   612.00   792.00",
		"Pages":     "1",
		"CropBox":   "0.00     0.00   612.00   792.00",
		"TrimBox":   "0.00     0.00   612.00   792.00",
		"ArtBox":    "0.00     0.00   612.00   792.00",
		"Optimized": "yes",
	}
	if !reflect.DeepEqual(meta, want_v018) {
		if !reflect.DeepEqual(meta, want) {
			t.Errorf("ConvertPDF() meta = %#v; not %#v", meta, want)
		}
	}
}
