FROM alpine

RUN apk update && apk add bash git

COPY tile-config-generator-linux /usr/bin/tile-config-generator
RUN chmod +x /usr/bin/tile-config-generator
