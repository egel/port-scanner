BINARY_NAME=port-scanner
BINARY_DIR=bin
WINDOWS=$(BINARY_DIR)/$(BINARY_NAME)_windows_amd64.exe
LINUX=$(BINARY_DIR)/$(BINARY_NAME)_linux_amd64
DARWIN=$(BINARY_DIR)/$(BINARY_NAME)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

build:
	env GOARCH=amd64 GOOS=darwin go build -v -o ${DARWIN} -ldflags="-s -w -X main.version=$(VERSION)" main.go
	env GOARCH=amd64 GOOS=linux go build -v -o ${LINUX} -ldflags="-s -w -X main.version=$(VERSION)" main.go
	env GOARCH=amd64 GOOS=windows go build -v -o ${WINDOWS} -ldflags="-s -w -X main.version=$(VERSION)" main.go

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

# used to force the specified targets to always execute,
# regardless of whether the output already exists
.PHONY: clean