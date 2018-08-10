FROM golang:1.9-alpine AS compilation

RUN apk update && apk add bash git unzip curl

ENV CGO_ENABLED=0
RUN go get -u github.com/pivotal-cf/om
RUN go get -u github.com/cloudfoundry/bosh-cli
FROM alpine

RUN apk update && apk add bash

COPY tile-config-generator-linux /usr/bin/tile-config-generator
COPY --from=compilation /go/bin/om /usr/bin
COPY --from=compilation /go/bin/bosh-cli /usr/bin/bosh
