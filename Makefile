all: help

test:
	@go test ./... -count=1 -cover

help:
	@echo 'Usage: '
	@echo 'make test'

.PHONY: all test help
