SHELL := /bin/bash

GO111MODULE = on

build:
	go build
.PHONY: build

test:
	go test -v
.PHONY: test

protobuf:
	protoc --go_out=. ./protobuf/*.proto
.PHONY: protobuf

mock-install:
	go get -u github.com/golang/mock/mockgen
	go get -u github.com/golang/mock
.PHONY: mock-install

mock:
	mockgen -package=internal -destination=client/internal/node_mock.go -source=client/internal/node.go
	mockgen -package=internal -destination=client/internal/store_mock.go -source=vendor/github.com/vacp2p/mvds/store/messagestore.go
.PHONY: mock

lint:
	golangci-lint run -v
.PHONY: lint

install-linter:
	# install linter
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.17.1
.PHONY: install-linter

vendor:
	go mod tidy
	go mod vendor
	modvendor -copy="**/*.c **/*.h" -v
.PHONY: vendor
