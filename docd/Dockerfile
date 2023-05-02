# Builder stage
FROM FROM --platform=linux/amd64 golang:1.20.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /usr/local/bin/docd

# Final stage
FROM FROM --platform=linux/amd64 debian:11

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

COPY --from=builder /usr/local/bin/docd /usr/local/bin/docd

RUN chmod a+x /usr/local/bin/docd

ENV PATH /usr/local/bin:$PATH

ENTRYPOINT ["/usr/local/bin/docd"]
