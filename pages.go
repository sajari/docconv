package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sajari/sajari-convert/iWork"
	"github.com/sajari/sajari-convert/snappy"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// Convert PAGES to text
func ConvertPages(input io.Reader) (string, map[string]string) {
	meta := make(map[string]string)
	var textBody string

	inputBytes, err := ioutil.ReadAll(input)
	if err != nil {
		log.Println("ioutil.ReadAll:", err)
		return "", nil
	}
	r, err := zip.NewReader(bytes.NewReader(inputBytes), int64(len(inputBytes)))
	if err != nil {
		log.Println("zip.NewReader:", err)
		return "", nil
	}

	for _, f := range r.File {
		if f.Name == "Index/Document.iwa" {
			rc, _ := f.Open()
			defer rc.Close()
			bReader := bufio.NewReader(snappy.NewReader(io.MultiReader(strings.NewReader("\xff\x06\x00\x00sNaPpY"), rc)))
			archiveLength, err := binary.ReadVarint(bReader)
			archiveInfoData, err := ioutil.ReadAll(io.LimitReader(bReader, archiveLength))
			archiveInfo := &TSP.ArchiveInfo{}
			err = proto.Unmarshal(archiveInfoData, archiveInfo)

			fmt.Println("archiveInfo:", archiveInfo, err)
		}
	}

	return textBody, meta
}
