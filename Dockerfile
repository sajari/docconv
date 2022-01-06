ARG alpine_version=3.14
FROM golang:1.17-alpine${alpine_version} as build

RUN apk add make \
            cmake \
            gcc \
            g++ \
            poppler-utils \
            wv \
            lynx \
            tesseract-ocr-dev

WORKDIR /go/src/code.sajari.com/docconv
COPY . /go/src/code.sajari.com/docconv

RUN go build -o dist/docconv -tags ocr code.sajari.com/docconv/docd

ARG alpine_version=3.14
FROM alpine:${alpine_version}

RUN apk add tesseract-ocr \
            tesseract-ocr-dev \
            poppler-utils

COPY --from=build /go/src/code.sajari.com/docconv/dist/docconv /usr/local/bindocconv

ENTRYPOINT ["/usr/local/bin/docconv"]
