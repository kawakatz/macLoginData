# Reference: https://github.com/projectdiscovery/httpx/blob/master/Makefile
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY=bin/macLoginData
GOBIN=~/go/bin
    
all: build
build:
		$(GOBUILD) -v -o "bin/macLoginData" cmd/macLoginData/macLoginData.go
test: 
		$(GOTEST) -v ./...
tidy:
		$(GOMOD) tidy
install:	$(BINARY)
		install -s $(BINARY) $(GOBIN)