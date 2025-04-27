FROM golang:1.23

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 go build -o main ./cmd/api

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 