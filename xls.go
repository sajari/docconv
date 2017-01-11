package docconv

import (
	"fmt"
	"github.com/extrame/xls"
	"io"
	"strings"
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
			row := w.GetSheet(i).Rows[uint16(j)]
			r := make([]string, 0)
			for _, col := range row.Cols {
				if uint16(len(r)) <= col.LastCol() {
					r = append(r, make([]string, col.LastCol()-uint16(len(r))+1)...)
				}
				str := col.String(w)
				for k := uint16(0); k < col.LastCol()-col.FirstCol()+1; k++ {
					r[col.FirstCol()+k] = str[k]
				}
			}

			body += strings.Join(r[:], ",")
			if j < int(w.GetSheet(i).MaxRow) {
				body += "\n"
			}
		}
	}

	return body, meta, nil
}
