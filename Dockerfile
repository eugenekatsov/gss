# Use an official Golang runtime as the base image
FROM golang:1.20.4-alpine3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the server will be listening on
EXPOSE 80

# Start the server
CMD ["./main"]
