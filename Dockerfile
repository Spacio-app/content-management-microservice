# Stage 1: Build the Go binary
FROM golang:1.24-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2: Run the binary
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

# Optionally expose port (should match what your app listens on)
EXPOSE 3000

CMD ["./main"]
