package docconv

import (
	"fmt"
	"io"

	"encoding/csv"

	"io/ioutil"

	"github.com/tealeg/xlsx"
)

// ConvertXLSX Excel Spreadsheet
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
	meta := make(map[string]string)
	meta["ModifiedDate"] = fmt.Sprintf("%d", fileStat.ModTime().Unix())

	// Document body
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

	return body, meta, nil
}

// CSVtoXLSX convert CSV data to XLSX
func CSVtoXLSX(r io.Reader) (xlsByte string, err error) {
	reader := csv.NewReader(r)

	var file *xlsx.File
	var sheet *xlsx.Sheet

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", fmt.Errorf("error try add sheet: %s", err.Error())
	}

	// write others rows
	for {
		rec, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		rows := sheet.AddRow()

		for _, h := range rec {
			row := rows.AddCell()
			row.Value = h
		}
	}

	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	err = file.Write(f)
	if err != nil {
		return "", fmt.Errorf("error try write xlsx: %s", err.Error())
	}

	xlsxB, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return
	}
	xlsByte = string(xlsxB)
	return
}
