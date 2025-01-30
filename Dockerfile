# Stage 1: Builder
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=arm64 go build -o main cmd/server/main.go

# Stage 2: Final image
FROM alpine:latest

# Install dependencies (wget and bash)
RUN apk add --no-cache wget bash

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main ./

# Copy migration files
COPY ./migrations /app/migrations

# Determine architecture and download the correct goose binary
ARG TARGETARCH
RUN case "$TARGETARCH" in \
      "amd64"|"") GOOSE_ARCH="amd64";; \
      "arm64") GOOSE_ARCH="arm64";; \
      *) echo "Unsupported architecture: $TARGETARCH; switching to amd64."; GOOSE_ARCH="amd64";; \
    esac && \
    echo "Requesting goose binary for $GOOSE_ARCH..." && \
    wget -O /bin/goose https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_${GOOSE_ARCH} && \
    chmod +x /bin/goose

# Add entrypoint script to handle migrations and run the server
COPY ./entrypoint.sh /app/entrypoint.sh

# Ensure entrypoint.sh has executable permissions
RUN chmod +x /app/entrypoint.sh

# Expose the application port
EXPOSE 8080

# Set the entrypoint
ENTRYPOINT ["/app/entrypoint.sh"]

# Start the application server
CMD ["./main"]