# Start from the official golang base image
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main .

# Start a new stage from scratch
FROM alpine:latest 

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /main

EXPOSE 3000

ENTRYPOINT [ "/main" ]
