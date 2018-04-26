FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/wabei/go-wabei && \
  (cd go- && make geth) && \
  cp go-/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go- && rm -rf /var/cache/apk/*

EXPOSE 8899
EXPOSE 17899

ENTRYPOINT ["/geth"]
