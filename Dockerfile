# Step 1: Use the official Go image to build the binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

# Step 2: Use a lightweight image to run the binary
FROM alpine:latest  

WORKDIR /app

# Copy the built binary from the builder stage
# (Docker will now correctly look at stage 1 instead of looking online)
COPY --from=builder /app/main .

# Expose the port your Go server listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]