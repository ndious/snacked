FROM golang:1.22-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=1

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    build-base \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

# RUN CGO_ENABLED=1 go install github.com/mattn/go-sqlite3@latest

EXPOSE 1337
EXPOSE 2345

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download && \
    go mod verify

RUN go build -o snacked -a .

### Production
FROM alpine:latest

RUN apk update && \
    apk add --no-cache \
        ca-certificates \
        curl \
        tzdata && \
    update-ca-certificates

# Copy executable
COPY --from=builder /app/snacked /usr/local/bin/snacked
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/snacked"]
