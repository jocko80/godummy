.DEFAULT_GOAL := check

.PHONY: fmt vet test race build check

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test:
	go test ./...

race:
	go test -race ./...

build: vet
	go build ./...

check: fmt vet test

