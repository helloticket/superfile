all: help

test:
	@go test ./... -count=1 -cover

coverage:
	@go get github.com/jstemmer/go-junit-report
	@go get github.com/axw/gocov/gocov
	@go get github.com/AlekSi/gocov-xml
	@go test -v -coverprofile=coverage.txt -covermode count ./... 2>&1 | go-junit-report > report.xml
	@gocov convert coverage.txt > coverage.json
	@gocov-xml < coverage.json > coverage.xml

help:
	@echo 'Usage: '
	@echo 'make test'

.PHONY: all test help
