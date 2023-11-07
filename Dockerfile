FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build
RUN make build

# set port
EXPOSE 8080

# Run
CMD ["./bin/aid"]