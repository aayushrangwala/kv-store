GOPATH = $(shell go env GOPATH)
GOOS=linux
GOARCH=amd64
GOOSTEST = $(shell go env GOOS)
VERSION ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git describe --always --dirty)
GIT_COMMIT ?= $(shell git rev-parse HEAD)
export GO111MODULE=on

all: build test

vendor:
	go mod tidy

build: fmt lint vet clean vendor
	CGO_ENABLED=0 go build -o bin/kv-store .

clean:
	rm -f bin/*

vet:
	go vet ./cmd/...
	go vet ./pkg/...
	go vet ./internal/...

lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.44.2
	$(GOPATH)/bin/golangci-lint run ./internal/...

test:
	GOOS=$(GOOSTEST) go test -count=1./internal/... -cover

fmt:
	gofmt -w ./internal
