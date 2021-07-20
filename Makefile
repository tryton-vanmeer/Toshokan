.PHONY: build clean

all: build

build:
	go build -o bin/toshokan

clean:
	rm -rf bin/