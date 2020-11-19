from golang:1.14-alpine

RUN apk add git && \
    go get -u github.com/rakyll/hey 

WORKDIR /$HOME

ENTRYPOINT ["hey","-c", "1000","-z","30s","-disable-redirects", "-m", "GET", "http://nginx"]

