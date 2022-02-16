.PHONY: build clean tool lint help
SHELL := /bin/bash
tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

run:
	go run main.go


test:
	./env.sh

help:
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make run: run go application"
	@echo "make test: test go application"

