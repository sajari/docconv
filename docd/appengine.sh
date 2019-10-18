#!/usr/bin/env sh

export NAME=docd
export VERSION=appengine

function die {
  echo 1>&2 $*
  exit 1
}

mkdir appengine/dependencies
cp alpine/dependencies/* appengine/dependencies

echo "Building ${NAME} for ${VERSION}..."

GOOS=linux GOARCH=amd64 go build -o $NAME || exit 1
cp docd $VERSION
cd $VERSION

echo "Deploying to AppEngine..."
gcloud app deploy --version="1"

# clean up
rm docd
rm -r appengine/dependencies