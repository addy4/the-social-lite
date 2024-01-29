# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Install dependencies using go modules (replace with your dependency management system if needed)
RUN go mod download

# Build the Go application
RUN go build services\main.go

# Expose ports 4010 and 8089
EXPOSE 4010
EXPOSE 8089

# Command to run the executable
CMD ["./main"]