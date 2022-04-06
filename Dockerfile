FROM golang:1.18-alpine3.15

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    apk add build-base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download -x && go mod verify

#Install wire
RUN go install github.com/google/wire/cmd/wire@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Run wire
RUN wire

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]