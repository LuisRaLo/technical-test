# Dockerfile para el backend en Go
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./test ./test

RUN GOOS=linux CGO_ENABLED=0 go build -o bootstrap ./cmd

EXPOSE 8080

RUN ls -la

CMD ["./bootstrap"]
