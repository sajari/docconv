#!/usr/bin/env sh

# Debian alternative build with OCR enabled.

# Build context must be the GOPATH where docconv and gosseract are contained. 

# Build runs on the Docker image, which is more reliable when working from other 
# OS than Linux.

export NAME=docd
export VERSION=debian
export DOCKERFILE=$GOPATH/src/code.sajari.com/docconv/docd/debian_ocr/Dockerfile

echo "Building ${NAME} for ${VERSION} with OCR enabled..."

echo "GOPATH: ${GOPATH}"

echo "Dockerfile: ${DOCKERFILE}"

docker build \
    -t $NAME:ocr \
    -f $DOCKERFILE \
    $GOPATH
