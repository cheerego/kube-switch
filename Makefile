SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.io

all:clean fmt
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/amd64/darwin/kube-switch -v  .
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/amd64/linux/kube-switch -v  .
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/amd64/windows/kube-switch.exe -v  .
clean:
	rm -rf dist
fmt:
	gofmt -w .

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"