GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=gofmt

.PHONY: all build fmt check
all: fmt check build

## Format source codes
fmt:
	$(GOFMT) -s -l -e -w .

## Check lint
check:
	go vet ./...
	golint src/...

## build binary
build:
	$(GOBUILD) -buildmode=plugin -o example.so example.go
