VERSION=0.1.0

LDFLAGS=-ldflags "-X \"github.com/mosuka/indigo/version.Version=${VERSION}\""

vendor:
	gvt restore

protoc:
	cd ${CURDIR}; protoc --go_out=plugins=grpc:. ./proto/indigo_service.proto

#test:
#	cd ${CURDIR}; go test

build:
	cd ${CURDIR}/indigo; go build ${LDFLAGS}
	cd ${CURDIR}/indigoctl; go build ${LDFLAGS}
	cd ${CURDIR}/indigorest; go build ${LDFLAGS}

install:
	cd ${CURDIR}/indigo; go install ${LDFLAGS}
	cd ${CURDIR}/indigoctl; go install ${LDFLAGS}
	cd ${CURDIR}/indigorest; go install ${LDFLAGS}

clean:
	cd ${CURDIR}/indigo; go clean
	cd ${CURDIR}/indigoctl; go clean
	cd ${CURDIR}/indigorest; go clean
