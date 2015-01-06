all: clean build test

.PHONY: all test clean

build:
	go install ./...

clean:
	go clean -i ./...

test:
	go test -v ./...
