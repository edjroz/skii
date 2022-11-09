#!/usr/bin/make -f

VERSION="0.0.1"
COMMIT=$(shell git log -1 --format='%H')

export GO111MODULE = on

verify: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

build: verify
	go build -mod=readonly $(BUILD_FLAGS) -o build/skii ./cmd/skiid

test: verify
	go test -mod=readonly 

lint: verify
	golangci-lint run
