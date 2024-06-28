# Dockerfile para el backend en Go
FROM golang:1.22 AS builder

WORKDIR /backend

# Instalar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Imagen final mínima
FROM alpine:latest

WORKDIR /app

# Copiar el ejecutable compilado
COPY --from=builder /backend/main .

# Comando por defecto para ejecutar el backend
CMD ["./main"]
