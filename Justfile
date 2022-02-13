@build:
	go build -o bin/toshokan main.go

@vendor:
	go mod vendor -v

@test:
	go test ./... -v --cover

@clean:
	rm -rf bin/