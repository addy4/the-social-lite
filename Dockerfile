FROM golang:latest

# set working directory
WORKDIR /app

# Copy the entire project
COPY . .

# working directory to the services directory
WORKDIR /app/services

# install dependencies
RUN go mod download

# build the app in the services directory
RUN go build -o main main.go

# Expose ports 4010 and 8089
EXPOSE 4010
EXPOSE 8089

# executable
CMD ["./main"]
