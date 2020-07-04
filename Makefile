GOCMD=go
GOBUILD=$(GOCMD) build

OUTPUT_NAME=brazen.exe
OUTPUT_DARWIN=brazen

build:
	$(GOBUILD) -o $(OUTPUT_NAME)

build-mac:
	$(GOBUILD) -o $(OUTPUT_DARWIN)