FROM debian

MAINTAINER Hamish Ogilvy

RUN apt-get update && apt-get install -y \
 zip \
 poppler-utils \
 wv \
 unrtf \
 tidy \
 lynx

# Port for Sajari
EXPOSE 8888

COPY docd /

# Start supervisor
CMD ["/docd"]
