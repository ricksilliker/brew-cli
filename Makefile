GOCMD=go
GOBIN=$(shell go env GOPATH)/bin
GOBUILD=$(GOCMD) build
GOINSTALL = $(GOCMD) install

OUTPUT_NAME=brazen.exe
OUTPUT_DARWIN=brazen

build:
	$(GOBUILD) -o $(OUTPUT_NAME)

build-mac:
	$(GOBUILD) -o $(OUTPUT_DARWIN)

install-mac:
	$(GOBUILD) -o $(GOBIN)/$(OUTPUT_DARWIN)
