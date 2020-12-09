.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o sec2env

.PHONY: build/linux
build/linux:
	env GOOS=linux GOARCH=amd64 go build -o sec2env
