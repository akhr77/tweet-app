FROM golang:1.13.1-alpine

ENV GO111MODULE=on

RUN apk update && \
    apk add vim curl git && \
    go get github.com/pilu/fresh
# イメージからコンテナを作成する際に実行
CMD ["fresh"]

