APP?=ekube
PORT?=7777
PROJECT?=github.com/bmsandoval/ekube
CONTAINER_IMAGE?=docker.io/bmsandoval/${APP}
DEV_IMAGE?=docker.io/bmsandoval/go-build

RELEASE?=0.0.3

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CURDIR?=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

GOOS?=linux
GOARCH?=amd64

push-staging:
	docker build -f ./deployment/Dockerfile-staging -t $(CONTAINER_IMAGE):$(RELEASE)-staging .
	docker push $(CONTAINER_IMAGE):$(RELEASE)-staging

local:
	helm upgrade --install dev-${APP} ./chart/ekube

remove:
	helm delete dev-${APP}

depend:
	go mod vendor

protoc:
	protoc -I servers/helloworld --go_out=plugins=grpc:servers/helloworld servers/helloworld/helloworld.proto

test:
	go test -v -race ./...

#.PHONY: charts
#all: charts
#
#charts:
#	cd chart && helm package ekube/
#	mv chart/*.tgz docs/
##	helm repo index docs --url https://alexellis.github.io/ekube/ --merge ./docs/index.yaml

