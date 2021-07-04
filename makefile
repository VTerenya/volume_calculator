format:
	go fmt volume_calculator/cmd

lint:
	golangci-lint run > linter.txt

test:
	go test ./...

build:
	go build volume_calculator/cmd

all:  format lint test build

