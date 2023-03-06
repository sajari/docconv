FROM golang:1.20 AS build

WORKDIR /app
COPY . ./
RUN go build -o /bin/docd ./docd

################################################################################

FROM debian:11-slim AS docd

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    ca-certificates \
    lynx \
    poppler-utils \
    tidy \
    unrtf \
    wv \
    zip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN update-ca-certificates

EXPOSE 8888

COPY --from=build /bin/docd /docd
ENTRYPOINT ["/docd"]
CMD ["--help"]
