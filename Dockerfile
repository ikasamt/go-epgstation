FROM golang:latest

WORKDIR /go/src

RUN go get github.com/cheekybits/genny
RUN go get github.com/ikasamt/zapp-jam
