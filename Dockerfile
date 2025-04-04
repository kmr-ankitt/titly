# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .
RUN go build -o /app/main .

# Final stage
FROM redis:latest

# Copy the built Go application from the builder stage
COPY --from=builder /app/main /app/main

# Expose ports for both the Go app and Redis
EXPOSE 4000 6379

# Start both Redis and the Go application
CMD ["sh", "-c", "redis-server & /app/main"]