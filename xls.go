package docconv

import (
	"fmt"
	"io"
	"strings"

	"github.com/extrame/xls"
)

// ConvertXLS Convert MS Excel Spreadsheet
func ConvertXLS(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	fileStat, err := f.Stat()
	if err != nil {
		return "", nil, fmt.Errorf("error on getting file stats: %v", err)
	}
	w, err := xls.OpenReader(f, "utf-8")
	if err != nil {
		return "", nil, fmt.Errorf("error on xls parsing: %v", err)
	}

	// Meta data
	meta := make(map[string]string)
	meta["ModifiedDate"] = fmt.Sprintf("%d", fileStat.ModTime().Unix())

	// Document body
	var body string
	for i := 0; i < w.NumSheets(); i++ {
		for j := 0; j <= int(w.GetSheet(i).MaxRow); j++ {
			row := w.GetSheet(i).Row(j)
			r := make([]string, 0)
			for k := row.FirstCol(); k == row.LastCol(); k++ {
				col := row.Col(i)
				r = append(r, col)
			}

			body += strings.Join(r[:], ",")
			if j < int(w.GetSheet(i).MaxRow) {
				body += "\n"
			}
		}
	}

	return body, meta, nil
}
