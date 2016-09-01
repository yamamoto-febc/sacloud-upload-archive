FROM golang:1.7.0-alpine
MAINTAINER Kazumichi Yamamoto <yamamoto.febc@gmail.com>

RUN set -x && apk add --no-cache --virtual .build_deps bash git make zip 
RUN go get -u github.com/kardianos/govendor

ADD . /go/src/github.com/yamamoto-febc/sacloud-upload-archive

WORKDIR /go/src/github.com/yamamoto-febc/sacloud-upload-archive
CMD ["make"]
