ORGANISATION = ldejager
VERSION ?= latest
COMPONENT = microbadger-notify

.PHONY: build

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

all: deps build build_docker

deps:
	go get -u github.com/ldejager/microbadger-notify

build: build_static build_cross build_tar build_sha

build_static:
	go install github.com/ldejager/microbadger-notify
	mkdir -p release
	cp $(GOPATH)/bin/microbadger-notify release/

build_cross:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o release/linux/amd64/microbadger-notify   github.com/ldejager/microbadger-notify
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o release/darwin/amd64/microbadger-notify  github.com/ldejager/microbadger-notify

build_tar:
	tar -cvzf release/linux/amd64/microbadger-notify.tar.gz   -C release/linux/amd64   microbadger-notify
	tar -cvzf release/darwin/amd64/microbadger-notify.tar.gz  -C release/darwin/amd64  microbadger-notify

build_sha:
	sha256sum release/linux/amd64/microbadger-notify.tar.gz   > release/linux/amd64/microbadger-notify.sha256
	sha256sum release/darwin/amd64/microbadger-notify.tar.gz  > release/darwin/amd64/microbadger-notify.sha256

build_docker:
	docker build --no-cache -t $(ORGANISATION)/$(COMPONENT):$(VERSION) .

default: all
