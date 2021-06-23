.PHONY: help clean fmt lint vet test test-cover generate-grpc build build-docker all

default: help

all:    ## clean, format, build and unit test
	make clean-all
	make gofmt
	make build
	make test

install:    ## build and install go application executable
	go install -v ./...


clean:  ## go clean
	rm -rf bin/

fmt:    ## format the go source files
	go fmt ./...

test:
	go test ./... -count=1

build:
	go build -o bin/handler ./cmd