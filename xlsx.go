package docconv

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
)

// Convert MS Excel Spreadsheet
func ConvertXLSX(r io.Reader) (string, map[string]string, error) {
	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	fileStat, err := f.Stat()
	if err != nil {
		return "", nil, fmt.Errorf("error on getting file stats: %v", err)
	}
	xlsFile, err := xlsx.OpenReaderAt(f, fileStat.Size())
	if err != nil {
		return "", nil, fmt.Errorf("error on xlsx parsing: %v", err)
	}

	// Meta data
	mc := make(chan map[string]string, 1)
	go func() {
		meta := make(map[string]string)
		meta["ModifiedDate"] = fmt.Sprintf("%d", fileStat.ModTime().Unix())
		mc <- meta
	}()

	// Document body
	bc := make(chan string, 1)
	go func() {
		var body string
		for _, sheet := range xlsFile.Sheets {
			for rowIndex, row := range sheet.Rows {
				for cellIndex, cell := range row.Cells {
					text, _ := cell.String()
					body += text
					if cellIndex < len(row.Cells)-1 {
						body += ","
					}
				}
				if rowIndex < len(sheet.Rows)-1 {
					body += "\n"
				}
			}
		}
		bc <- body
	}()

	body := <-bc
	meta := <-mc

	return body, meta, nil
}
