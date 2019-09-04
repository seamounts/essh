ARCHS ?= amd64
GOOS ?= linux

BINARY=essh
VERSION=`git describe --tags`
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"


build:
	GOOS=$(GOOS) GOARCH=$(ARCHS) go build ${LDFLAGS} -o ${BINARY}

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
