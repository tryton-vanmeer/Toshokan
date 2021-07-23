.PHONY: build clean

all: build

build:
	go build -o bin/toshokan src/main.go

clean:
	rm -rf bin/