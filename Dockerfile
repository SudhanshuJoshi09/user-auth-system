# Use the official Golang image from Docker Hub
FROM golang AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new image from the official Golang image
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/main .
COPY --from=build /app/.env ./

# Expose the port the app will run on
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]

