# --- Build stage ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Copy the rest of the rest-api code
COPY . .

# Build the REST API
RUN go build -mod=vendor -o taskape-rest-api .

# --- Final stage ---
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/taskape-rest-api .

# Expose port
EXPOSE 8080

# Run the REST API
CMD ["./taskape-rest-api"]
