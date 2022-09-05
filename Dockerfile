FROM docker.io/golang:1.14-buster AS build
RUN apt update && \
    apt install -y software-properties-common && \
    apt-get --assume-yes install apt-transport-https && \
    apt-add-repository 'deb https://notesalexp.org/tesseract-ocr-dev/buster/ buster main' && \
    apt-get update -oAcquire::AllowInsecureRepositories=true && \
    apt-get --assume-yes --allow-unauthenticated install notesalexp-keyring -oAcquire::AllowInsecureRepositories=true && \
    apt-get update && \
    apt-get --assume-yes install tesseract-ocr libtesseract-dev tesseract-ocr-tha
WORKDIR /go/src/code.sajari.com/docconv/
COPY ./ ./
WORKDIR /go/src/code.sajari.com/docconv/docd/
RUN go get -t github.com/otiai10/gosseract/...
RUN go build -tags ocr .
FROM docker.io/debian:11-slim AS docd
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        lynx \
        poppler-utils \
        tidy \
        unrtf \
        wv \
        zip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
RUN apt update && \
    apt install -y software-properties-common && \
    apt-get --assume-yes install apt-transport-https && \
    apt-add-repository 'deb https://notesalexp.org/tesseract-ocr-dev/bullseye/ bullseye main' && \
    apt-get update -oAcquire::AllowInsecureRepositories=true && \
    apt-get --assume-yes --allow-unauthenticated install notesalexp-keyring -oAcquire::AllowInsecureRepositories=true && \
    apt-get update && \
    apt-get --assume-yes install tesseract-ocr libtesseract-dev tesseract-ocr-tha
EXPOSE 8888
COPY --from=build /go/src/code.sajari.com/docconv/docd/docd /docd
ENTRYPOINT ["/docd"]
CMD ["--help"]