SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'

PLATFORM := $(shell uname -o)

NAME := Simple.exe
OS := windows

ifeq ($(PLATFORM), Msys)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else ifeq ($(PLATFORM), Cygwin)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else
	INCLUDE := $(GOPATH)
	NAME=Simple
	OS=linux
endif

# enable second expansion
.SECONDEXPANSION:

.PHONY: all
.PHONY: pbs
.PHONY: test

BINDIR=$(INCLUDE)/bin

all: pbs sol build

build:
	GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o $(BINDIR)/$(NAME)

pbs:
	cd pbs/ && $(MAKE)

sol:
	cd contract/ && $(MAKE)

resdir := ./webserver
webfs:
	go-bindata -o $(resdir)/webfs/webfs.go -pkg=webfs $(resdir)/html/dist/...

mac:
	GOOS=darwin go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).mac
arm:
	GOOS=linux GOARM=7 GOARCH=arm go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm
lnx:
	GOOS=linux go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).lnx

clean:
	rm $(BINDIR)/$(NAME)
