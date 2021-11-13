FROM golang:1.17.3-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download
