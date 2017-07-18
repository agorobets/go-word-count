export PATH := $(GOPATH)/bin:$(PATH)

NOW := $(shell date -u +%Y%m%d.%H%M%S)
GITCOMMIT := $(shell git describe --always)
LDFLAGS := -X main.VERSION=$(NOW)-$(GITCOMMIT)-dev -X main.GITCOMMIT=$(GITCOMMIT) -X main.BUILT=$(NOW)

all: install

install:
	go install -v -ldflags "${LDFLAGS}"

run: install
	go-word-count

.PHONY: doc
doc:
	godoc -http=":6060"
