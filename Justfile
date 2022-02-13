@build:
	go build -o bin/toshokan src/main.go

@vendor:
	go mod vendor -v

@test:
	go test ./src/util

@clean:
	rm -rf bin/