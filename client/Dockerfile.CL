# Use an official Golang runtime as a parent image
FROM golang:1.21.2

# Set the working directory in the container
WORKDIR /app

# Copy the client.go file into the container
COPY client.go /app/
COPY go.mod/ /app/
COPY go.sum/ /app/
COPY templates/ /app/templates
COPY pb/ /app/pb


# Build the Go application inside the container
RUN go build client.go

# Expose the port for your web server
EXPOSE 8081

# Command to run the Go client
CMD ["./client"]
