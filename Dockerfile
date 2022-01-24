# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal

RUN go build -o main ./cmd/main

CMD [ "./main" ]
