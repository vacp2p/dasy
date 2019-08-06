SHELL := /bin/bash

GO111MODULE = on

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
