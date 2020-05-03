FROM golang:1.13-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
RUN mkdir -p /app
WORKDIR /app
ADD . /app
ADD .env.example /app/.env

COPY . .

# Build the Go app
RUN go build .


# This container exposes port 8080 to the outside world
EXPOSE 8070

# Run the binary program produced by `go install`
ENTRYPOINT ["./ziswaf-backend"]