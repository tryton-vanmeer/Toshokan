@build:
	go build -o bin/toshokan src/main.go

@test:
	go test ./src/util

@clean:
	rm -rf bin/