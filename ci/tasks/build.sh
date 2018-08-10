#!/bin/bash -e

mkdir ~/.ssh/ && touch ~/.ssh/known_hosts
ssh-keyscan github.com >>~/.ssh/known_hosts

export GOPATH=$PWD/go
export PATH=$GOPATH/bin:$PATH
OUTPUT_DIR=$PWD/compiled-output
SOURCE_DIR=$PWD/source

cp source/Dockerfile ${OUTPUT_DIR}/.

go get github.com/Masterminds/glide
go get github.com/xchapter7x/versioning

cd ${SOURCE_DIR}
if [ -d ".git" ]; then
  DRAFT_VERSION=`versioning bump_patch`-`git rev-parse HEAD`
else
  DRAFT_VERSION="v0.0.0-local"
fi
echo "next version should be: ${DRAFT_VERSION}"

WORKING_DIR=$GOPATH/src/github.com/pivotalservices/tile-config-generator
mkdir -p ${WORKING_DIR}
cp -R ${SOURCE_DIR}/* ${WORKING_DIR}/.
cd ${WORKING_DIR}
glide install
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUTPUT_DIR}/tile-config-generator-linux -ldflags "-X github.com/pivotalservices/tile-config-generator/commands.VERSION=${DRAFT_VERSION}" cmd/tile-config-generator/main.go
GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT_DIR}/tile-config-generator-osx -ldflags "-X github.com/pivotalservices/tile-config-generator/commands.VERSION=${DRAFT_VERSION}" cmd/tile-config-generator/main.go
GOOS=windows GOARCH=amd64 go build -o ${OUTPUT_DIR}/tile-config-generator.exe -ldflags "-X github.com/pivotalservices/tile-config-generator/commands.VERSION=${DRAFT_VERSION}" cmd/tile-config-generator/main.go

echo ${DRAFT_VERSION} > ${OUTPUT_DIR}/name
echo ${DRAFT_VERSION} > ${OUTPUT_DIR}/tag
