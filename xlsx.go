package docconv

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize"
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

	xlsFile, err := excelize.OpenFile(f.Name())
	if err != nil {
		fmt.Println(".... > ", err)
		return "", nil, fmt.Errorf("error on xlsx parsing: %v", err)
	}

	// Meta data
	meta := make(map[string]string)
	meta["ModifiedDate"] = fmt.Sprintf("%d", fileStat.ModTime().Unix())

	// Document body
	var body string
	for sheetIndex, _ := range xlsFile.GetSheetMap() {
		rows := xlsFile.GetRows("sheet" + strconv.Itoa(sheetIndex))
		for rowIndex, row := range rows {
			for cellIndex, colCell := range row {
				text := colCell
				body += text
				if cellIndex < len(row)-1 {
					body += ","
				}
			}
			if rowIndex < len(rows)-1 {
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
