NAME=mock
REVISION:=$(shell git rev-parse --short HEAD)
LDFLAGS := -X main.revision=${REVISION}

GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=gofmt
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGENERATE=$(GOCMD) generate

BINARY=$(NAME)

.PHONY: all test build clean check fmt generate
all: generate fmt check test build

## Show help
help:
	@make2help $(MAKEFILE_LIST)

## Format source codes
fmt:
	$(GOFMT) -s -l -e -w .

## Check lint
check:
	errcheck -exclude errcheck_excludes.txt -asserts -verbose ./...
	go vet ./...
	golint src/...

## GO test
test:
	$(GOTEST) -v ./...

## build binary
build:
	$(GOBUILD) -o bin/$(BINARY) -ldflags "$(LDFLAGS)" src/main.go

## Generate assets files
generate:
	$(GOGENERATE) ./...

## rm binary
clean:
	$(GOCLEAN)
	rm -fr bin/
