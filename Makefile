# Variables
MAIN_PKG := ./cmd/releases
LDFLAGS  := -ldflags "-s -w"

# Default mode
.DEFAULT_GOAL := build

build:
	go build $(LDFLAGS) $(MAIN_PKG)

test:
	go test -v $(MAIN_PKG)

.PHONY: build test