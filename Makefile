build:
	GOPATH=$(shell pwd) go build -o bin/generate ./cmd/generate \
	go build -o bin/slicer ./cmd/slicer
