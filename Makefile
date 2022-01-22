.PHONY: build
build:
	go build -o build/mcargo -v ./cmd/mcargo/mcargo.go

.PHONY: test
test:
	go test -v -race -timeout 30s .

.PHONY: clean
clean:
	rm -rf build

.DEFAULT_GOAL := build