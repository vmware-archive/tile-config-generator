FROM golang:1.9-alpine AS compilation

RUN apk update && apk add bash git unzip curl

ENV CGO_ENABLED=0
COPY . $GOPATH/src/github.com/pivotalservices/tile-config-generator
RUN go get -a -t github.com/pivotalservices/tile-config-generator/...
RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get -u github.com/golang/dep/cmd/dep
RUN cd $GOPATH/src/github.com/pivotalservices/tile-config-generator && dep ensure && ginkgo -r
RUN go build github.com/pivotalservices/tile-config-generator/cmd/tile-config-generator
RUN go get -a github.com/pivotal-cf/om
RUN go get -a github.com/cloudfoundry/bosh-cli
FROM alpine

RUN apk update && apk add bash

COPY --from=compilation /go/bin/tile-config-generator /usr/bin
COPY --from=compilation /go/bin/om /usr/bin
COPY --from=compilation /go/bin/bosh-cli /usr/bin/bosh
