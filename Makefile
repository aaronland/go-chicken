CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep
	if test -d src/github.com/thisisaaronland/go-chicken; then rm -rf src/github.com/thisisaaronland/go-chicken; fi
	mkdir -p src/github.com/thisisaaronland/go-chicken
	cp -r strings src/github.com/thisisaaronland/go-chicken/
	cp -r emoji src/github.com/thisisaaronland/go-chicken/
	cp chicken.go src/github.com/thisisaaronland/go-chicken/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/cooperhewitt/go-ucd"
	@GOPATH=$(GOPATH) go get -u "github.com/facebookgo/grace/gracehttp"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-sanitize"

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/chicken cmd/chicken.go
	@GOPATH=$(GOPATH) go build -o bin/rooster cmd/rooster.go

fmt:
	@GOPATH=$(GOPATH) go fmt cmd/*.go
	@GOPATH=$(GOPATH) go fmt chicken.go
	@GOPATH=$(GOPATH) go fmt strings/*.go
	@GOPATH=$(GOPATH) go fmt emoji/*.go

alpha:
	@GOPATH=$(GOPATH) go run cmd/build-alpha-codes.go > emoji/emoji.go

docker-build:
	docker build -t rooster .

docker-debug: docker-build
	docker run -it -p 1280:1280 rooster