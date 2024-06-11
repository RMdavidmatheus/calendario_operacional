# Stage 1: Build the Go application
FROM golang:1.22 AS builder

WORKDIR /usr/src/app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o app ./cmd/main.go

# Stage 2: Create a lightweight image for running the application
FROM debian:stable-slim

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /usr/src/app/app .

CMD ["./app"]