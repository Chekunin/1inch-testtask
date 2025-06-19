FROM golang:1.23.4-alpine AS builder

COPY . /app/
WORKDIR /app

RUN apk add --no-cache gcc musl-dev mailcap

RUN go mod download && go build -o 1inch_testtask ./cmd/

ENTRYPOINT [ "./1inch_testtask" ]
