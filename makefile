t:
	go test ./...

b:
	go build ./cmd

f:
	go fmt ./...

l:
	golangci-lint run ./... > lint.txt
all: f l b t