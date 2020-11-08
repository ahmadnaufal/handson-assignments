TIMESTAMP = $(shell date +'%Y%m%d%H%M%S')
PWD	      = $(shell pwd)

IMAGE_NAME = rentuffdocker/currency-converter
VERSION    = $(shell git show -q --format=%h)

# LOCAL SETUP FOR COMPILED BINARY
GOOS	 ?= linux
GOARCH  = amd64
ODIR    = _output

export GO111MODULE ?= on

all: compile test

test:
	go test -cover ./...

coverage:
	go tool cover -html=cover.out -o cover.html

bin:
	go build -o currency-converter app/main.go

mod:
	go mod tidy

compile:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $(ODIR)/currency-converter app/main.go

build:
	docker build -t $(IMAGE_NAME):$(VERSION) -f ./Dockerfile .
