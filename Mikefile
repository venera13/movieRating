all: build test check

modules:
	go mod tidy

build: modules
	go build -v -o bin/ratingservice cmd/ratingservice/*.go

test:
	go test ./...

check:
	golangci-lint run