FROM golang:alpine AS base

WORKDIR /usr/src/app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    build-base \
    && update-ca-certificates

##
# Development with hot reload and debugger
###
FROM base AS dev
WORKDIR /usr/src/app

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./internal ./internal
COPY ./cmd ./cmd
COPY ./docker/go/.air.toml ./.air.toml
RUN go build -o . ./cmd/www

EXPOSE 1337
EXPOSE 2345

ENTRYPOINT ["air"]

##
# Test runner config
###
FROM base AS test-watch

WORKDIR /usr/src/app

RUN go install github.com/meian/gowatch@latest

COPY ./go.mod ./go.sum ./
RUN go mod download && \
    go mod verify

ENTRYPOINT ["gowatch", "test"]
