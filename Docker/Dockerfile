# Start from a Debian-based image with Go installed as a build stage
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /project

# Copy the go.mod file and download dependencies
COPY ./go.mod .
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./httpserver

# Start a new stage from scratch
FROM debian:buster-slim

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /project/main /app/

# Copy the static files from the project
COPY ./public /app/public

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program
CMD ["/app/main"]