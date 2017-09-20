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
	// TODO: find out why the order isn't being respected
	var body string
	for i := 0; i < w.NumSheets(); i++ {
		for _, row := range w.GetSheet(i).Rows {
			r := make([]string, 0)
			for _, j := range row.Cols {
				if uint16(len(r)) <= j.LastCol() {
					r = append(r, make([]string, j.LastCol()-uint16(len(r))+1)...)
				}
				str := j.String(w)
				for k := uint16(0); k < j.LastCol()-j.FirstCol()+1; k++ {
					r[j.FirstCol()+k] = str[k]
				}
			}

			body += strings.Join(r[:], ",")
			body += "\n"
		}
	}

	return body, meta, nil
}
