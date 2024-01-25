# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.21 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN make build

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:stable-slim

WORKDIR /app
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/bin/aid /app/bin/aid
COPY config.toml /app/config.toml

ENV GO_ENV=release

EXPOSE 8080

# Run the web service on container startup.
CMD ["bin/aid"]


