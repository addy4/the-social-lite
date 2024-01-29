# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Change the working directory to the services directory
WORKDIR /app/services

# Install dependencies using go modules
RUN go mod download

# Build the Go application in the services directory
RUN go build -o main main.go

# Expose ports 4010 and 8089
EXPOSE 4010
EXPOSE 8089

# Command to run the executable
CMD ["./main"]
