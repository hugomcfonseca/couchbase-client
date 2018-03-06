FROM golang:alpine3.7 as builder

LABEL maintainer "Hugo Fonseca <https://github.com/hugomcfonseca>"

WORKDIR /go/src/github.com/hugomcfonseca/couchbase-tasker/app

RUN \
    apk add --update --no-cache git \
        && go get -d -v \
        && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cb_tasker .

FROM alpine:3.7

LABEL maintainer "Hugo Fonseca <https://github.com/hugomcfonseca>"

ENV \
    CB_ACTION='init'

COPY --from=builder /go/src/github.com/hugomcfonseca/couchbase-tasker/app/cb_tasker /usr/local/bin

ENTRYPOINT [ "cb_tasker" ]
