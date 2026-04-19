BINARY     := harry-potter
ALIASES    := hp harry potter
MODULE     := codeberg.org/zatarain/harry-potter
VERSION    := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT     := $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS    := -ldflags "-X $(MODULE)/command.Version=$(VERSION) -X $(MODULE)/command.Commit=$(COMMIT) -X $(MODULE)/command.BuildDate=$(BUILD_DATE)"
DESTDIR    ?= /usr/local/bin

.PHONY: build clean install uninstall lint vet fmt check-fmt test coverage tidy

build:
	go build $(LDFLAGS) -o bin/$(BINARY) .

tidy:
	go mod tidy

clean:
	rm -rf bin/

install: build
	install -m 755 bin/$(BINARY) $(DESTDIR)/$(BINARY)
	for alias in $(ALIASES); do \
		ln -sf $(DESTDIR)/$(BINARY) $(DESTDIR)/$$alias; \
	done

uninstall:
	rm -f $(DESTDIR)/$(BINARY)
	for alias in $(ALIASES); do \
		rm -f $(DESTDIR)/$$alias; \
	done

lint:
	golangci-lint run ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

check-fmt:
	@test -z "$$(gofmt -l .)" || { gofmt -l .; exit 1; }

test:
	go test -race -count=1 ./...

coverage:
	go test -race -count=1 -coverprofile=coverage.out -covermode=atomic ./...

.DEFAULT_GOAL := build
