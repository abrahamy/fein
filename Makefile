.DEFAULT_GOAL := build

lint:
	golangci-lint run ./...
.PHONY:lint

build: lint
	echo "done"
.PHONY:build