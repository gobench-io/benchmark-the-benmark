from golang:1.14-alpine

RUN apk add git && \
    go get -u github.com/rakyll/hey 

WORKDIR /$HOME

# Max RPS
# ENTRYPOINT ["hey","-c", "1000","-z","30s","-disable-redirects", "-m", "GET", "http://nginx"] 
# MEM/VU
# ENTRYPOINT ["hey","-c", "1","-z","10s","-disable-redirects", "-m", "GET", "http://nginx"]
ENTRYPOINT ["hey","-c", "30","-z","120s","-disable-redirects", "-m", "GET", "http://172.31.3.99"]
# ENTRYPOINT ["hey","-c", "10","-z","10s","-disable-redirects", "-m", "GET", "http://nginx"]
# ENTRYPOINT ["hey","-c", "20","-z","10s","-disable-redirects", "-m", "GET", "http://nginx"]

