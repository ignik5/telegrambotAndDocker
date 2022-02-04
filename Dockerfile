FROM golang:1.17 AS builder
ENV GO111MODULE=on

WORKDIR /build
#go modules
COPY go.mod go.sum ./ 

RUN go mod download
COPY ./src ./src

RUN cd ./src && go build -ldflags "-s -w" -o wrapper

