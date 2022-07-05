FROM golang:1.18-alpine

RUN mkdir -p /usr/share/man/man1

RUN apk add --no-cache ca-certificates tzdata bash git alpine-sdk automake make autoconf unzip gcc gnupg gdal-dev curl wget tidyhtml lynx wv poppler-utils

# build unrtf
WORKDIR /unrtf
RUN wget https://ftp.gnu.org/gnu/unrtf/unrtf-0.21.10.tar.gz && \
  tar xzvf unrtf-0.21.10.tar.gz && \
  cd unrtf-0.21.10/ && \
  ./bootstrap && \
  ./configure && \
  make && \
  make install

WORKDIR /docconv

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . ./

RUN go test -race ./...