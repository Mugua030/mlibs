APPNAME=mlibs
PWD=$(shell pwd)
GOOS=$(shell go env |grep "GOOS=" | tr -d '"' | awk -F "=" '{print $$2}')

all:
	@echo GOOS=$(GOOS)
	env CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -o ./cmd/mlibs/${appName} ./cmd/mlibs/

run:
	./cmd/mlibs/mlibs
