FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=arm64 go build -o main cmd/server/main.go

FROM alpine:latest

RUN apk add --no-cache make git go bash

WORKDIR /app

COPY --from=builder /app/main ./

EXPOSE 8080

CMD ["./main"]