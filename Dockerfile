FROM golang:1.24-bookworm

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy
