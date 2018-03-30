# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=pattern

all: build
build:
	$(GOBUILD) -v
	$(GOBUILD) -v ./cmd/pattern
	cd ../..
install:	
	$(GOINSTALL) -v
	$(GOINSTALL) -v ./cmd/pattern
	cd ../..

