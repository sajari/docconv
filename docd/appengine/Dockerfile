FROM alpine

MAINTAINER Hamish Ogilvy
 
ENV CC=/usr/bin/gcc
ENV CXX=/usr/bin/g++

COPY dependencies/* /

RUN apk add --update --no-cache make cmake gcc g++ poppler-utils=0.48.0-r0 wv=1.2.9-r3 lynx=2.8.8_p2-r5; ./install.sh; apk del make cmake gcc g++

EXPOSE 8080

COPY docd /
CMD ["/docd", "-addr", ":8080"]
