GOCMD=go
GOBUILD=$(GOCMD) build

OUTPUT_NAME=brazen.exe

build:
	$(GOBUILD) -o $(OUTPUT_NAME)