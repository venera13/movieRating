all: build test check

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down

modules:
	go mod tidy

build: modules
	go build -v -o bin/ratingservice cmd/ratingservice/*.go

test:
	go test ./...

check:
	golangci-lint run