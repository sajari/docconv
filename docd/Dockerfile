FROM golang:1.19 AS build

WORKDIR /app
COPY . ./
RUN go build ./docd/...

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

COPY --from=build /app/docd /docd
ENTRYPOINT ["/docd"]
CMD ["--help"]
