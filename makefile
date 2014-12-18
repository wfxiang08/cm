all: clean build

build:
	go install ./...

clean:
	go clean -i ./...

integration_test:
	sh run_test.sh
