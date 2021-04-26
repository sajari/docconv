package html_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"

	"code.sajari.com/docconv"
)

func TestConvertHTML_readabilityUseClasses(t *testing.T) {
	tests := []struct {
		readabilityUseClasses string
		want                  string
	}{
		// The default value set by this package.
		{
			readabilityUseClasses: "",
			want:                  ``,
		},
		// Enable output with readability true.
		{
			readabilityUseClasses: "good",
			want: `1
word
This is a full sentence.
`,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q", tt.readabilityUseClasses), func(t *testing.T) {
			old := docconv.HTMLReadabilityOptionsValues.ReadabilityUseClasses
			t.Cleanup(func() {
				docconv.HTMLReadabilityOptionsValues.ReadabilityUseClasses = old
			})
			docconv.HTMLReadabilityOptionsValues.ReadabilityUseClasses = tt.readabilityUseClasses

			data, err := ioutil.ReadFile("testdata/test.html")
			must(t, err)
			got, _, err := docconv.ConvertHTML(bytes.NewReader(data), true)
			must(t, err)

			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("result mismatch (-want +got):\n%v", diff)
			}
		})
	}
}

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
