package docconv

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/extrame/xls"
)

// ConvertXLS Convert MS Excel Spreadsheet
func ConvertXLS(r io.Reader) (string, map[string]string, error) {
	f, err := ioutil.TempFile("", "nuveo")
	if err != nil {
		return "", nil, fmt.Errorf("error on create temp file: %v", err)
	}
	defer os.Remove(f.Name())

	var byt []byte
	byt, err = ioutil.ReadAll(r)
	if err != nil {
		return "", nil, fmt.Errorf("error on readAll: %v", err)
	}
	_, err = f.Write(byt)
	if err != nil {
		return "", nil, fmt.Errorf("error on write temp file: %v", err)
	}

	_ = f.Close()

	w, err := xls.Open(f.Name(), "utf-8")
	if err != nil {
		return "", nil, fmt.Errorf("error on xls parsing: %v", err)
	}
	// Meta data
	meta := make(map[string]string)
	meta["ModifiedDate"] = fmt.Sprintf("%d", time.Now().Unix())
	// Document body
	var body string
	for i := 0; i < w.NumSheets(); i++ {
		for j := 0; j <= int(w.GetSheet(i).MaxRow); j++ {
			row := w.GetSheet(i).Row(j)
			r := make([]string, 0)
			for k := row.FirstCol(); k <= row.LastCol(); k++ {
				col := row.Col(k)
				r = append(r, col)
			}
			body += strings.Join(r[:], ",")
			if j < int(w.GetSheet(i).MaxRow) {
				body += "\n"
			}
		}
	}
	return strings.TrimSuffix(body, "\n"), meta, nil
}
