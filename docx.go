package docconv

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

type TypeOverride struct {
	XMLName     xml.Name `xml:"Override"`
	ContentType string   `xml:"ContentType,attr"`
	PartName    string   `xml:"PartName,attr"`
}

type Type struct {
	XMLName   xml.Name       `xml:"Types"`
	Overrides []TypeOverride `xml:"Override"`
}

// ConvertDocx converts an MS Word docx file to text.
func ConvertDocx(r io.Reader) (string, map[string]string, error) {
	meta := make(map[string]string)
	var textHeader, textBody, textFooter string

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil, err
	}
	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return "", nil, fmt.Errorf("error unzipping data: %v", err)
	}

	zipFiles := mapZipFiles(zr.File)

	types, err := getContentTypes(zipFiles["[Content_Types].xml"])
	if err != nil {
		return "", nil, err
	}

	for _, override := range types.Overrides {
		f := zipFiles[override.PartName]

		switch {
		case override.ContentType == "application/vnd.openxmlformats-package.core-properties+xml":
			rc, err := f.Open()
			if err != nil {
				return "", nil, fmt.Errorf("error opening '%v' from archive: %v", f.Name, err)
			}
			defer rc.Close()

			meta, err = XMLToMap(rc)
			if err != nil {
				return "", nil, fmt.Errorf("error parsing '%v': %v", f.Name, err)
			}

			if tmp, ok := meta["modified"]; ok {
				if t, err := time.Parse(time.RFC3339, tmp); err == nil {
					meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
				}
			}
			if tmp, ok := meta["created"]; ok {
				if t, err := time.Parse(time.RFC3339, tmp); err == nil {
					meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
				}
			}
		case override.ContentType == "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml":
			body, err := parseDocxText(f)
			if err != nil {
				return "", nil, err
			}
			textBody += body + "\n"
		case override.ContentType == "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml":
			footer, err := parseDocxText(f)
			if err != nil {
				return "", nil, err
			}
			textFooter += footer + "\n"
		case override.ContentType == "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml":
			header, err := parseDocxText(f)
			if err != nil {
				return "", nil, err
			}
			textHeader += header + "\n"
		}

	}
	return textHeader + "\n" + textBody + "\n" + textFooter, meta, nil
}

func getContentTypes(f *zip.File) (*Type, error) {
	contentTypesFile, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer contentTypesFile.Close()

	contentTypesFileBytes, err := ioutil.ReadAll(contentTypesFile)
	if err != nil {
		return nil, err
	}

	var types Type
	err = xml.Unmarshal(contentTypesFileBytes, &types)
	if err != nil {
		return nil, err
	}
	return &types, nil
}

func mapZipFiles(files []*zip.File) map[string]*zip.File {
	filesMap := map[string]*zip.File{}
	for _, f := range files {
		filesMap[f.Name] = f
		filesMap["/"+f.Name] = f
	}
	return filesMap
}

func parseDocxText(f *zip.File) (string, error) {
	r, err := f.Open()
	if err != nil {
		return "", fmt.Errorf("error opening '%v' from archive: %v", f.Name, err)
	}
	defer r.Close()

	text, err := DocxXMLToText(r)
	if err != nil {
		return "", fmt.Errorf("error parsing '%v': %v", f.Name, err)
	}
	return text, nil
}

// DocxXMLToText converts Docx XML into plain text.
func DocxXMLToText(r io.Reader) (string, error) {
	return XMLToText(r, []string{"br", "p", "tab"}, []string{"instrText", "script"}, true)
}
