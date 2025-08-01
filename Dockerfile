# Build stage
FROM golang:1.24-bookworm AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .

# Puerto por defecto
EXPOSE 3001

# Comando de inicio
CMD ["./main"]