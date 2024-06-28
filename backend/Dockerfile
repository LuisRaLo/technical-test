# Dockerfile para el backend en Go
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY ./go.mod ./

RUN go mod download

COPY ./cmd ./cmd

RUN GOOS=linux CGO_ENABLED=0 go build -o bootstrap ./cmd

EXPOSE 8000

CMD ["./backend"]
