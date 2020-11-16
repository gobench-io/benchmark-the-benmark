from golang:1.14-alpine

RUN apk add git && \
    go get -u github.com/rakyll/hey 

WORKDIR /$HOME

ENTRYPOINT ["hey", "-n", "1000","-c", "1000", "-m", "GET", "http://nginx"]

