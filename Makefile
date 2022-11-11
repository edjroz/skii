#!/usr/bin/make -f

export GO111MODULE = on

verify: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

build: verify
	go build -mod=readonly -o build/skii ./cmd/skiid

install: verify
	go install -mod=readonly ./cmd/skiid

verify: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

test: verify
	go test -mod=readonly 

lint: verify
	golangci-lint run
