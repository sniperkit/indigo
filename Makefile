VERSION=0.1.0

LDFLAGS=-ldflags "-X \"github.com/mosuka/indigo/version.Ver=${VERSION}\""

vendor:
	gvt restore

protoc:
	cd ${CURDIR}; protoc --go_out=plugins=grpc:. ./proto/indigo_service.proto

#test:
#	cd ${CURDIR}; go test

build:
	cd ${CURDIR}; go build ${LDFLAGS}

install:
	cd ${CURDIR}; go install ${LDFLAGS}

clean:
	cd ${CURDIR}; go clean
