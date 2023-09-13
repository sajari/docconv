package docconv

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestConvertDoc(t *testing.T) {
	if _, err := exec.LookPath("wvText"); err != nil {
		t.Skip("wvText not installed")
		return
	}

	tests := []struct {
		file            string
		wantTrimmedText string
		wantMeta        map[string]string
		wantErr         bool
	}{
		{
			file:            "001-test.doc",
			wantTrimmedText: "test",
			wantMeta: map[string]string{
				"AppName":            "Microsoft Office Word",
				"CharCount":          "4",
				"Character count":    "4",
				"CodePage":           "1252",
				"Company":            "",
				"CreateTime":         "2023-09-13 01:54:00 +0000 UTC",
				"CreatedDate":        "1694570040",
				"Dirty links":        "false",
				"DocSecurity":        "0",
				"Document parts":     "0",
				"EditTime":           "1970-01-01 00:00:00 +0000 UTC",
				"Heading pair":       "0",
				"Hyperlinks changed": "false",
				"LastAuthor":         "cloudconvert_7",
				"LastSaveTime":       "2023-09-13 01:54:00 +0000 UTC",
				"Line count":         "1",
				"ModifiedDate":       "1694570040",
				"PageCount":          "1",
				"Paragraph count":    "1",
				"RevNumber":          "1",
				"Scale":              "false",
				"Shared document":    "false",
				"Template":           "Normal",
				"Version":            "1048576",
				"WordCount":          "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			f, err := os.Open(path.Join("testdata", tt.file))
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			gotText, gotMeta, err := ConvertDoc(f)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotText = strings.TrimSpace(gotText)
			if gotText != tt.wantTrimmedText {
				t.Errorf("ConvertDoc() text = %v, want %v", gotText, tt.wantTrimmedText)
			}
			if !cmp.Equal(tt.wantMeta, gotMeta, maybeTimeComparer) {
				t.Errorf("ConvertDoc() meta mismatch (-want +got):\n%v", cmp.Diff(tt.wantMeta, gotMeta, maybeTimeComparer))
			}
		})
	}
}

// Compares strings as time.Times if they look like times. Required because
// wvText returns different time formats depending on system clock.
var maybeTimeComparer = cmp.Comparer(func(x, y string) bool {
	xt, xterr := time.Parse("2006-01-02 15:04:05 -0700 MST", x)
	yt, yterr := time.Parse("2006-01-02 15:04:05 -0700 MST", y)
	if xterr == nil && yterr == nil {
		return xt.Equal(yt)
	}
	return x == y
})
