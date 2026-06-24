FROM golang:alpine

WORKDIR /app

# Copy dependency files and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy all your source code into the container
COPY . .

# Build the app directly
RUN go build -o /workdirk cmd/server/main.go

# Expose the port your server listens on
EXPOSE 8080

# Run the binary directly
CMD ["/workdirk"]