# Use the official Golang image as a build stage
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o devops-hello-world .

# Use a smaller base image for the final stage
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/devops-hello-world .

# Ensure the binary has execution permissions
RUN chmod +x /app/devops-hello-world

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["/app/devops-hello-world"]
