CGO_ENABLED=1
.PHONY: build
build:
	go build -v ./cmd/server

.PHONY: test
test:
	go test -v -timeout 30s ./...

.DEFAULT_GOAL := build