FROM golang:1.13.1-alpine

ENV GO111MODULE=on

COPY . /go/src

WORKDIR /go/src/app

RUN apk update && \
    apk add vim curl git && \
    go get github.com/pilu/fresh
# go mod download

CMD ["fresh"]
