package docconv

import (
	"strings"
	"testing"
)

func TestNewLocalFile(t *testing.T) {
	_, err := NewLocalFile(strings.NewReader("test"))
	if err != nil {
		t.Fatalf("got error %v, want nil", err)
	}
}
