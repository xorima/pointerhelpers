.PHONY: all build test clean

all: test build

build:
	@echo "Building the project..."
	go build github.com/xorima/pointerhelpers

test:
	@echo "Running tests..."
	go test ./... --race

clean:
	@echo "Cleaning up..."
	go clean
