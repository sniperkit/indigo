#  Copyright (c) 2017 Minoru Osuka
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# 		http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION=0.1.0

LDFLAGS=-ldflags "-X \"github.com/mosuka/indigo/version.Version=${VERSION}\""

vendoring:
	gvt restore

protoc:
	cd ${CURDIR}; protoc --go_out=plugins=grpc:. ./proto/indigo_service.proto

#test:
#	cd ${CURDIR}; go test

build:
	cd ${CURDIR}/indigo; go build -tags=lang ${LDFLAGS}
	cd ${CURDIR}/indigoctl; go build -tags=lang ${LDFLAGS}
	cd ${CURDIR}/indigorest; go build -tags=lang ${LDFLAGS}

install:
	cd ${CURDIR}/indigo; go install -tags=lang ${LDFLAGS}
	cd ${CURDIR}/indigoctl; go install -tags=lang ${LDFLAGS}
	cd ${CURDIR}/indigorest; go install -tags=lang ${LDFLAGS}

clean:
	cd ${CURDIR}/indigo; go clean
	cd ${CURDIR}/indigoctl; go clean
	cd ${CURDIR}/indigorest; go clean
