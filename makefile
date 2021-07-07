.PHONY: test

format:
	go fmt ./...

build:
	go build ./cmd

lint:
	golangci-lint run -c ./.golangci.yml > lint.txt

test:
	go test ./...

all:  format lint build test
