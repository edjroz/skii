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

generate-protoc: verify
	protoc --go_out=./types/ --go_opt=paths=source_relative \
    --go-grpc_out=./types/ --go-grpc_opt=paths=source_relative \
    proto/skii.proto
