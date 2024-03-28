FROM debian

RUN apt-get update 
RUN apt-get install -y zip
RUN apt-get install -y poppler-utils
RUN apt-get install -y wv
RUN apt-get install -y unrtf
RUN apt-get install -y tidy
RUN apt-get install -y lynx
RUN apt-get install -y libtesseract-dev
RUN apt-get install -y libleptonica-dev
RUN apt-get install -y tesseract-ocr-eng
RUN apt-get install -y git
RUN apt-get install -y golang

# Build context must be the host GOPATH
COPY . /goworkspace

WORKDIR /goworkspace/src/code.sajari.com/docconv/docd

ENV GOPATH=/goworkspace

RUN GOOS=linux GOARCH=amd64 go build -tags ocr -o /docd

EXPOSE 8888

CMD ["/docd"]
