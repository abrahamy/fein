.DEFAULT_GOAL := build

lint:
	go mod tidy
	golangci-lint run ./...
.PHONY:lint

test: lint
	go test ./...
.PHONY:test

build: lint
	echo "done"
.PHONY:build