CWD=$(shell pwd)
GOPATH := $(CWD)/vendor:$(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep
	if test -d src/github.com/thisisaaronland/go-chicken; then rm -rf src/github.com/thisisaaronland/go-chicken; fi
	mkdir -p src/github.com/thisisaaronland/go-chicken
	cp -r strings src/github.com/thisisaaronland/go-chicken/
	cp chicken.go src/github.com/thisisaaronland/go-chicken/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps bin

deps:
	if test ! -d vendor; then mkdir -p vendor; fi
	@GOPATH=$(GOPATH) go get -u "github.com/cooperhewitt/go-ucd"
	find vendor -name '.git' -print -type d -exec rm -rf {} +

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/chicken cmd/chicken.go

fmt:
	go fmt cmd/*.go
	go fmt chicken.go
	go fmt strings/*.go
