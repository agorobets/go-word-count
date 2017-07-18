export PATH := $(GOPATH)/bin:$(PATH)

all: install

install:
	go install 

run: install
	go-word-count

.PHONY: doc
doc:
	godoc -http=":6060"
