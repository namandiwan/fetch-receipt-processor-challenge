# Use the official Go image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o receipt-processor

# Expose port 8080 for the API
EXPOSE 8080

# Command to run the application
CMD ["./receipt-processor"]