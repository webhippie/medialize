.PHONY: clean deps test build

# export GOOS ?= linux
# export GOARCH ?= amd64
export CGO_ENABLED ?= 0

LDFLAGS += -X "main.buildDate=$(shell date -u '+%Y-%m-%d %H:%M:%S %Z')"

clean:
	go clean -i ./...

deps:
	go get -t ./...

test:
	go test -cover ./...

build:
	go build -ldflags '-s -w $(LDFLAGS)'
