FROM golang:latest

ENV GOOS=darwin
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/src
