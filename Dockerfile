FROM golang:1.12-alpine3.9 as builder

RUN apk add --nocache --update alpine-sdk

COPY . /go/src/gitlab.com/target-smart-data-ai-search/pg-log-extractor

RUN go build -o /go/bin/extractor  /go/src/gitlab.com/target-smart-data-ai-search/pg-log-extractor/cmd

FROM alpine

ENV "GOPATH" "/go"

RUN apk add --update tzdata ca-certificates openssl && rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/extractor /usr/local/bin/

WORKDIR /
