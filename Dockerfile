# Start from the official Golang image for building
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o chain-healthCheck ./server/main.go

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/chain-healthCheck .

# Copy any config files if needed
COPY config ./config

# Expose port (change if your app uses a different port)
EXPOSE 3002

# Run the binary
CMD ["./chain-healthCheck"]