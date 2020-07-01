// +build ocr

package docconv

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestConvertImage(t *testing.T) {
	tests := []struct {
		image    string
		wantText string
		wantMeta map[string]string
		wantErr  bool
	}{
		{
			image:    "001-helloworld.png",
			wantText: "Hello, World!",
			wantMeta: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.image, func(t *testing.T) {
			image, err := os.Open(path.Join("testdata", tt.image))
			if err != nil {
				t.Fatal(err)
			}
			defer image.Close()

			gotText, gotMeta, err := ConvertImage(image)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotText != tt.wantText {
				t.Errorf("ConvertImage() text = %v, want %v", gotText, tt.wantText)
			}
			if !reflect.DeepEqual(gotMeta, tt.wantMeta) {
				t.Errorf("ConvertImage() meta = %v, want %v", gotMeta, tt.wantMeta)
			}
		})
	}
}
