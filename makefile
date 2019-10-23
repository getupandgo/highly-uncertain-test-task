SHELL := /bin/bash

all: deps

deps: go.mod go.sum
	@go mod download

build-container:
	docker build .