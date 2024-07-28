
.PHONY: lint build

build:
	go build -o bin/ ./cmd/...
lint:
	golangci-lint run 