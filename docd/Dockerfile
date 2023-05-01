FROM golang:1.20 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/app

################################################################################

FROM debian:11-slim 

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

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

ENV PATH /usr/local/bin:$PATH

CMD ["app"]