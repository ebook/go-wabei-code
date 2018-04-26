.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/wabei/go-wabei/cmd/swarm
	go build -o bin/geth     github.com/wabei/go-wabei/cmd/geth
	go build -o bin/bootnode github.com/wabei/go-wabei/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/wabei/go-wabei/swarm/...
