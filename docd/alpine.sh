#!/usr/bin/env sh

export NAME=docd
export VERSION=alpine

function die {
  echo 1>&2 $*
  exit 1
}

echo "Building ${NAME} for ${VERSION}..."

GOOS=linux GOARCH=amd64 go build -o $NAME || exit 1
cp docd $VERSION
cd $VERSION
docker build -t $NAME . || exit 1
rm docd

