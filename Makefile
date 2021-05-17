SHELL := /bin/bash

.EXPORT_ALL_VARIABLES:
SRC_DIR := $(shell pwd)
OUT_DIR := $(SRC_DIR)/_output
BIN_DIR := $(OUT_DIR)/bin
CONFIG_LOCATION := $(SRC_DIR)
GO111MODULE := on

include .env
export $(shell sed 's/=.*//' .env)

.PHONY: launch
launch:
	go run $(SRC_DIR)/main.go launch

.PHONY: migrations
migrations:
	go run $(SRC_DIR)/main.go migrations

.PHONY: rollbacks
rollbacks:
	go run $(SRC_DIR)/main.go rollbacks

.PHONY: build
build:
	go build -ldflags="-s -w" -o $(BIN_DIR)/server $(SRC_DIR)/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: test.nocache
test.nocache:
	go clean -testcache
	go test -v ./...

.PHONY: compose.up
compose.up:
	docker-compose up -d

.PHONY: compose.down
compose.down:
	docker-compose stop
	docker-compose rm -f
	docker-compose down -v

.PHONY: swag.install
swag.install:
	which swag || GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag

.PHONY: swag.generate
swag.generate:
	swag init -g internal/server/router.go