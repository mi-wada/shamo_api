FROM golang:1.16.5-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download
