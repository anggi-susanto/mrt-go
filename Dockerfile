# Build stage
FROM golang:1.21-alpine AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o ./bin/app ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the built executable from the build stage
COPY --from=build /app/bin .

# Expose the application port
EXPOSE 3000

# Run the Go application
CMD ["./app"]