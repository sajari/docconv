#!/bin/sh

VERSION=$1

apk add make cmake gcc g++ poppler-utils wv lynx tesseract-ocr-dev
go build -o dist/docconv-${VERSION} -tags ocr github.com/skpr/docconv/docd
