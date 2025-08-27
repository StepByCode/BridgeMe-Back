# syntax=docker/dockerfile:1

# Use the official Golang image to create a build artifact.
# This is known as a multi-stage build.
FROM golang:1.25-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# -o /app/main: output the binary to /app/main
# CGO_ENABLED=0: disable CGO for static linking
# GOOS=linux: build for linux
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Start a new stage from scratch for a smaller image
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
