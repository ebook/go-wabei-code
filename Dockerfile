# Build Geth in a stock Go builder container
FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-
RUN cd /go- && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-/build/bin/geth /usr/local/bin/

EXPOSE 8899 8546 17899 17899/udp 30304/udp
ENTRYPOINT ["geth"]
