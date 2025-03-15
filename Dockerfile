# Use the official Golang image
FROM golang:1.24.1 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for caching dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

COPY .env .env

# Move to the cmd directory to build the app
WORKDIR /app/cmd

# ✅ Fix build command
RUN go build -o main .

# Set working directory back to /app/cmd to run the binary
WORKDIR /app

# Expose the application port
EXPOSE 8082

# ✅ Fix execution path
CMD ["./cmd/main"]
