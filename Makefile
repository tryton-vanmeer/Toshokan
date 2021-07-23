.PHONY: build clean

all: build

build:
	go build -o bin/toshokan src/main.go

test:
	go test

clean:
	rm -rf bin/