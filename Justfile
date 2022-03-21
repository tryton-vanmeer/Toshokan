@build:
	go build -o bin/toshokan main.go

@run:
	go run main.go

@vendor:
	go mod vendor -v

@tidy:
	go mod tidy

@test:
	go test ./... -v --cover

@clean:
	rm -rf bin/