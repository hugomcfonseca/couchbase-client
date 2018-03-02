FROM golang:alpine3.7 as builder

# Build Go application to after copy to final container
WORKDIR /go/src/github.com/alexellis/href-counter/

RUN go get -d -v

COPY app.go    .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.7

LABEL maintainer "Hugo Fonseca <https://github.com/hugomcfonseca>"

ENV \
    CB_ACTION='init'

WORKDIR /app

COPY --from=builder /go/path /app

ENTRYPOINT [ "" ]

CMD
