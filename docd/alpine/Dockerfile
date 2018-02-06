FROM alpine:3.7

MAINTAINER Hamish Ogilvy
 
ENV CC=/usr/bin/gcc
ENV CXX=/usr/bin/g++

COPY dependencies/* /

RUN apk add --update --no-cache make cmake gcc g++ poppler-utils=0.56.0-r0 wv=1.2.9-r3 lynx=2.8.8_p2-r6 \
    && ./install.sh \
    && apk del make cmake gcc g++

EXPOSE 8888

COPY docd /
CMD ["/docd"]
