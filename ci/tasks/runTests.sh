#!/bin/bash -e

export GOPATH=$PWD/go
export PATH=$GOPATH/bin:$PATH

go get -u github.com/golang/dep/cmd/dep
WORKING_DIR=$GOPATH/src/github.com/pivotalservices/tile-config-generator
mkdir -p ${WORKING_DIR}
cp -R source/* ${WORKING_DIR}/.
cd ${WORKING_DIR}
go version
dep ensure
go test ./generator -v
