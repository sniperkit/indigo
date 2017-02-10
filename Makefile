VERSION=0.1.0

LDFLAGS=-ldflags "-X \"github.com/mosuka/indigo/version.Version=${VERSION}\""

vendor:
	gvt restore

protoc:
	cd ${CURDIR}; protoc --go_out=plugins=grpc:. ./proto/indigo_service.proto

#test:
#	cd ${CURDIR}; go test

build:
	cd ${CURDIR}/indigo_cli; go build ${LDFLAGS}
	cd ${CURDIR}/indigo_grpc; go build ${LDFLAGS}
	cd ${CURDIR}/indigo_rest; go build ${LDFLAGS}

install:
	cd ${CURDIR}/indigo_cli; go install ${LDFLAGS}
	cd ${CURDIR}/indigo_grpc; go install ${LDFLAGS}
	cd ${CURDIR}/indigo_rest; go install ${LDFLAGS}

clean:
	cd ${CURDIR}/indigo_cli; go clean
	cd ${CURDIR}/indigo_grpc; go clean
	cd ${CURDIR}/indigo_rest; go clean
