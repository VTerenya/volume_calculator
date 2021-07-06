f:
	go fmt ./...

b:
	go build ./cmd

l:
	golangci-lint run ./... > lint.txt

t:
	go test ./...

all:  f l b t