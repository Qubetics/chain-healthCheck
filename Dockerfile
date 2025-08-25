# Start from the official Golang image for building
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o echo-healthCheck ./server/main.go

# Use a minimal image for running
FROM alpine:latest

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/echo-healthCheck .

# Copy any config files if needed
COPY config ./config
# COPY .env .env

# Expose port (change if your app uses a different port)
EXPOSE 3002

# Run the binary
CMD ["./echo-healthCheck"]


# docker build -t echo-healthcheck . && docker run -d --env-file .env -p 3002:3002 echo-healthcheck