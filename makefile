all: clean build

build:
	go install ./...

clean:
	go clean -i ./...
