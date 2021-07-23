.PHONY: build lint test clean

all: build

build:
	@go build -o bin/toshokan src/main.go

lint:
	@golint ./...

test:
	@go test ./src/util

clean:
	rm -rf bin/