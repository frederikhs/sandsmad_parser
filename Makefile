all: download vet lint test

download:
	go mod download

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test -race -cover -v ./...
