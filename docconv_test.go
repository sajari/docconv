package docconv

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name        string
		r           io.Reader
		mimeType    string
		readability bool
		resp        *Response
		err         error
	}{
		{
			name:     "trims space",
			mimeType: "text/plain",
			r:        strings.NewReader(" \n\n\nthe file\n\n"),
			resp: &Response{
				Body: "the file",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.r, tt.mimeType, tt.readability)
			if err != tt.err {
				t.Fatalf("Convert() error = %v, want %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.resp) {
				t.Errorf("Convert() = %v, want %v", got, tt.resp)
			}
		})
	}
}
