# base image for the app
FROM golang:1.19
# instructs Docker to use this directory as the default destination for all subsequent commands
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# name of the binary will be the last arg
RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-app
EXPOSE 8080
CMD ["/chat-app"]