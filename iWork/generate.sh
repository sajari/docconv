#!/usr/bin/env bash
# Generate code from protos stored in pb-schema.

set -euo pipefail

cd "$(dirname "$0")"

function ensure_installed() {
  command -v "$1" >/dev/null 2>&1 || {
    echo >&2 "Could not find $1, please ensure it is installed."
    echo "Try running: $2"
    exit 1
  }
}

ensure_installed protoc "brew install protobuf"
ensure_installed protoc-gen-go "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"

protoc -I=./pb-schema --go_out=paths=source_relative:. TSPArchiveMessages.proto
protoc -I=./pb-schema --go_out=paths=source_relative:. TSPDatabaseMessages.proto
protoc -I=./pb-schema --go_out=paths=source_relative:. TSPMessages.proto
