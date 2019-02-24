build:
	GOPATH=$(shell pwd) go build -o bin/generate ./cmd/generate \
	&& GOPATH=$(shell pwd) go build -o bin/slicer ./cmd/slicer
