.PHONY: build
build:
	go build -o build/mcargo -v ./cmd/mcargo

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build