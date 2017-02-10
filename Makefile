VERSION=0.1.0

LDFLAGS=-ldflags "-X \"github.com/mosuka/bleve-server/version.Version=${VERSION}\""

vendor:
	gvt restore

protoc:
	#cd ${CURDIR}; protoc --go_out=plugins=grpc:. ./proto/bleve_service.proto
	cd ${CURDIR}; protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. ./proto/bleve_service.proto
	cd ${CURDIR}; protoc -I/usr/local/include -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/bleve_service.proto

#test:
#	cd ${CURDIR}; go test

build:
	cd ${CURDIR}/bleve-cli; go build ${LDFLAGS}
	cd ${CURDIR}/bleve-server; go build ${LDFLAGS}
	cd ${CURDIR}/bleve-rest; go build ${LDFLAGS}

install:
	cd ${CURDIR}/bleve-cli; go install ${LDFLAGS}
	cd ${CURDIR}/bleve-server; go install ${LDFLAGS}
	cd ${CURDIR}/bleve-rest; go install ${LDFLAGS}

clean:
	cd ${CURDIR}/bleve-cli; go clean
	cd ${CURDIR}/bleve-server; go clean
	cd ${CURDIR}/bleve-rest; go clean
