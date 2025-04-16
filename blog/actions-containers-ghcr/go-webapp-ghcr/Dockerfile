# ---- Build Stage ----
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

# Copy go mod file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN GOOS=linux go build -o app main.go

# ---- Run Stage ----
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary and template files from the builder stage
COPY --from=builder /app/app /app/index.tmpl ./

# Expose port
EXPOSE 8080

# Run the application
CMD ["./app"]