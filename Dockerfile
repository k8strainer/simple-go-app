# Use official Golang image as builder
FROM docker.io/library/golang:1.18 AS builder

WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o simple-go-app main.go

# Use a smaller image for running the application
FROM alpine:3.14

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/simple-go-app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./simple-go-app"]
