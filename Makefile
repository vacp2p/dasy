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
.PHONY: mock
