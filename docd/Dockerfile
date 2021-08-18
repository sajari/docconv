FROM golang:1.14 AS build

WORKDIR /go/src/github.com/sajari/docconv/docd
COPY main.go .
RUN go get -v ./
RUN go build .

################################################################################

FROM debian:11-slim AS docd

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

EXPOSE 8888

COPY --from=build /go/bin/docd /docd
ENTRYPOINT ["/docd"]
CMD ["--help"]

