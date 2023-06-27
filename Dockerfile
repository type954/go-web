# Start from a Go language base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod ./

# Download and cache Go module dependencies
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the command to run the executable
CMD ["./main"]

